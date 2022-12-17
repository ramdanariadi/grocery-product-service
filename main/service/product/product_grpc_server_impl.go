package product

import (
	"database/sql"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/helpers"
	"github.com/ramdanariadi/grocery-product-service/main/models"
	categoryRepo "github.com/ramdanariadi/grocery-product-service/main/repositories/category"
	"github.com/ramdanariadi/grocery-product-service/main/repositories/product"
	"github.com/ramdanariadi/grocery-product-service/main/service/category"
	"github.com/ramdanariadi/grocery-product-service/main/service/response"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"golang.org/x/net/context"
	"time"
)

type ProductServiceServerImpl struct {
	Repository  product.ProductRepositoryImpl
	RedisClient *redis.Client
}

func NewProductServiceServerImpl(db *sql.DB) *ProductServiceServerImpl {
	return &ProductServiceServerImpl{
		Repository: product.ProductRepositoryImpl{
			DB: db,
		},
		RedisClient: utils.NewRedisClient(),
	}
}

func (server ProductServiceServerImpl) FindById(ctx context.Context, id *ProductId) (*ProductResponse, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	var productModel *models.ProductModel
	cache, _ := server.RedisClient.Get(ctx, id.GetId()).Result()
	helpers.LogIfError(err)

	if cache != "" {
		err := json.Unmarshal([]byte(cache), productModel)
		helpers.LogIfError(err)
	} else {
		productModel = server.Repository.FindById(ctx, tx, id.Id)
		bytes, err := json.Marshal(productModel)
		helpers.LogIfError(err)
		err = server.RedisClient.Set(ctx, id.GetId(), bytes, 1*time.Hour).Err()
		helpers.LogIfError(err)
	}

	grpcProductModel := Product{
		Id:          productModel.Id,
		Name:        productModel.Name,
		Weight:      uint32(productModel.Weight),
		Category:    productModel.Category,
		ImageUrl:    productModel.ImageUrl,
		CategoryId:  productModel.CategoryId,
		Price:       productModel.Price,
		PerUnit:     uint64(productModel.PerUnit),
		Description: productModel.Description,
	}
	return &ProductResponse{
		Message: "OK",
		Status:  "Success",
		Data:    &grpcProductModel,
	}, nil
}

func (server ProductServiceServerImpl) FindProductsByCategory(ctx context.Context, id *category.CategoryId) (*MultipleProductResponse, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	if err != nil {
		return nil, err
	}
	raws := server.Repository.FindByCategory(ctx, tx, id.Id)
	products := fetchProducts(raws)
	status, message := utils.ResponseForQuerying(len(products) > 0)
	return &MultipleProductResponse{
		Status:  status,
		Data:    products,
		Message: message,
	}, nil
}

func (server ProductServiceServerImpl) FindRecommendedProduct(ctx context.Context, _ *ProductEmpty) (*MultipleProductResponse, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	if err != nil {
		return nil, err
	}
	raws := server.Repository.FindWhere(ctx, tx, "products.is_recommended = $1", true)
	products := fetchProducts(raws)
	status, message := utils.ResponseForQuerying(len(products) > 0)
	return &MultipleProductResponse{
		Status:  status,
		Data:    products,
		Message: message,
	}, nil
}

func (server ProductServiceServerImpl) FindTopProducts(ctx context.Context, _ *ProductEmpty) (*MultipleProductResponse, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	if err != nil {
		return nil, err
	}
	raws := server.Repository.FindWhere(ctx, tx, "products.is_top = $1", true)
	products := fetchProducts(raws)
	status, message := utils.ResponseForQuerying(len(products) > 0)
	return &MultipleProductResponse{
		Status:  status,
		Data:    products,
		Message: message,
	}, nil
}

func (server ProductServiceServerImpl) FindAll(ctx context.Context, _ *ProductEmpty) (*MultipleProductResponse, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	raws := server.Repository.FindAll(ctx, tx)
	products := fetchProducts(raws)
	status, message := utils.ResponseForQuerying(len(products) > 0)
	return &MultipleProductResponse{
		Status:  status,
		Data:    products,
		Message: message,
	}, nil
}

func fetchProducts(rows *sql.Rows) []*Product {
	var products []*Product
	for rows.Next() {
		productTmp := Product{ImageUrl: ""}
		var imageUrl sql.NullString
		err := rows.Scan(&productTmp.Id, &productTmp.Name, &productTmp.Price, &productTmp.PerUnit,
			&productTmp.Weight, &productTmp.Category, &productTmp.CategoryId,
			&productTmp.Description, &imageUrl)
		helpers.PanicIfError(err)

		if imageUrl.Valid {
			productTmp.ImageUrl = imageUrl.String
		}

		products = append(products, &productTmp)
	}
	return products
}

func (server ProductServiceServerImpl) Save(ctx context.Context, product *Product) (*response.Response, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	categoryRepository := categoryRepo.NewCategoryRepository(server.Repository.DB)
	categoryModel := categoryRepository.FindById(ctx, tx, product.CategoryId)
	if utils.IsTypeEmpty(categoryModel) {
		status, _ := utils.ResponseForQuerying(false)
		return &response.Response{Status: status, Message: "INVALID_CATEGORY"}, nil
	}

	id, _ := uuid.NewUUID()
	productModel := models.ProductModel{
		Id:          id.String(),
		Name:        product.Name,
		Weight:      uint(product.Weight),
		Category:    categoryModel.Category,
		CategoryId:  categoryModel.Id,
		Price:       product.Price,
		PerUnit:     uint(product.PerUnit),
		Description: product.Description,
		ImageUrl:    product.ImageUrl,
	}
	saved := server.Repository.Save(ctx, tx, productModel)
	status, message := utils.ResponseForModifying(saved)
	return &response.Response{
		Status:  status,
		Message: message,
	}, nil
}

func (server ProductServiceServerImpl) Update(ctx context.Context, product *Product) (*response.Response, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	categoryRepository := categoryRepo.NewCategoryRepository(server.Repository.DB)
	categoryModel := categoryRepository.FindById(ctx, tx, product.CategoryId)
	if utils.IsTypeEmpty(categoryModel) {
		status, _ := utils.ResponseForQuerying(false)
		return &response.Response{Status: status, Message: "INVALID_CATEGORY"}, nil
	}

	productModel := models.ProductModel{
		Id:          product.Id,
		Name:        product.Name,
		Weight:      uint(product.Weight),
		Category:    categoryModel.Category,
		CategoryId:  categoryModel.Id,
		Price:       product.Price,
		PerUnit:     uint(product.PerUnit),
		Description: product.Description,
		ImageUrl:    product.ImageUrl,
	}
	updated := server.Repository.Update(ctx, tx, productModel)
	status, message := utils.ResponseForModifying(updated)
	return &response.Response{
		Status:  status,
		Message: message,
	}, nil
}

func (server ProductServiceServerImpl) Delete(ctx context.Context, id *ProductId) (*response.Response, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	deleted := server.Repository.Delete(ctx, tx, id.Id)
	status, message := utils.ResponseForModifying(deleted)
	defer helpers.CommitOrRollback(tx)
	return &response.Response{Status: status, Message: message}, nil
}

func (server ProductServiceServerImpl) mustEmbedUnimplementedProductServiceServer() {
	//TODO implement me
	panic("implement me")
}

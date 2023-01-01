package product

import (
	"database/sql"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/category"
	repository2 "github.com/ramdanariadi/grocery-product-service/main/category/repository"
	"github.com/ramdanariadi/grocery-product-service/main/helpers"
	"github.com/ramdanariadi/grocery-product-service/main/product/model"
	"github.com/ramdanariadi/grocery-product-service/main/product/repository"
	"github.com/ramdanariadi/grocery-product-service/main/response"
	"github.com/ramdanariadi/grocery-product-service/main/setup"
	"golang.org/x/net/context"
	"time"
)

type ProductServiceServerImpl struct {
	Repository  repository.ProductRepositoryImpl
	RedisClient *redis.Client
}

func NewProductServiceServerImpl(db *sql.DB) *ProductServiceServerImpl {
	return &ProductServiceServerImpl{
		Repository: repository.ProductRepositoryImpl{
			DB: db,
		},
		RedisClient: setup.NewRedisClient(),
	}
}

func (server ProductServiceServerImpl) FindById(ctx context.Context, id *ProductId) (*ProductResponse, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	var productModel *model.ProductModel
	cache, err := server.RedisClient.Get(ctx, id.GetId()).Result()
	helpers.LogIfError(err)

	if cache != "" {
		err := json.Unmarshal([]byte(cache), productModel)
		helpers.LogIfError(err)
	} else {
		productModel = server.Repository.FindById(ctx, tx, id.Id)
		if productModel == nil {
			return &ProductResponse{
				Message: "EMPTY",
				Status:  "FAILED",
				Data:    nil,
			}, nil
		}
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

	rows := server.Repository.FindByCategory(ctx, tx, id.Id)
	products := fetchProducts(rows)
	status, message := setup.ResponseForQuerying(len(products) > 0)
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

	rows := server.Repository.FindWhere(ctx, tx, "products.is_recommended = $1", true)
	products := fetchProducts(rows)
	status, message := setup.ResponseForQuerying(len(products) > 0)
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

	rows := server.Repository.FindWhere(ctx, tx, "products.is_top = $1", true)
	products := fetchProducts(rows)
	status, message := setup.ResponseForQuerying(len(products) > 0)
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

	rows := server.Repository.FindAll(ctx, tx)
	products := fetchProducts(rows)
	status, message := setup.ResponseForQuerying(len(products) > 0)
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
	helpers.LogIfError(rows.Close())
	return products
}

func (server ProductServiceServerImpl) Save(ctx context.Context, product *Product) (*response.Response, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	categoryRepository := repository2.NewCategoryRepository(server.Repository.DB)
	categoryModel := categoryRepository.FindById(ctx, tx, product.CategoryId)
	if categoryModel != nil {
		status, _ := setup.ResponseForQuerying(false)
		return &response.Response{Status: status, Message: "INVALID_CATEGORY"}, nil
	}

	id, _ := uuid.NewUUID()
	productModel := model.ProductModel{
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
	err = server.Repository.Save(ctx, tx, &productModel)
	status, message := setup.ResponseForModifying(err == nil)
	return &response.Response{
		Status:  status,
		Message: message,
	}, nil
}

func (server ProductServiceServerImpl) Update(ctx context.Context, product *Product) (*response.Response, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	categoryRepository := repository2.NewCategoryRepository(server.Repository.DB)
	categoryModel := categoryRepository.FindById(ctx, tx, product.CategoryId)
	if categoryModel != nil {
		status, _ := setup.ResponseForQuerying(false)
		return &response.Response{Status: status, Message: "INVALID_CATEGORY"}, nil
	}

	productModel := model.ProductModel{
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
	err = server.Repository.Update(ctx, tx, &productModel)
	status, message := setup.ResponseForModifying(err == nil)
	return &response.Response{
		Status:  status,
		Message: message,
	}, nil
}

func (server ProductServiceServerImpl) Delete(ctx context.Context, id *ProductId) (*response.Response, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	err = server.Repository.Delete(ctx, tx, id.Id)
	status, message := setup.ResponseForModifying(err == nil)
	return &response.Response{Status: status, Message: message}, nil
}

func (server ProductServiceServerImpl) mustEmbedUnimplementedProductServiceServer() {
	//TODO implement me
	panic("implement me")
}

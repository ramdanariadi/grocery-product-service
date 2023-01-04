package product

import (
	"database/sql"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/category"
	repository2 "github.com/ramdanariadi/grocery-product-service/main/category/repository"
	"github.com/ramdanariadi/grocery-product-service/main/product/model"
	"github.com/ramdanariadi/grocery-product-service/main/product/repository"
	"github.com/ramdanariadi/grocery-product-service/main/response"
	"github.com/ramdanariadi/grocery-product-service/main/setup"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"golang.org/x/net/context"
	"log"
	"time"
)

type ProductServiceServerImpl struct {
	Repository  repository.ProductRepository
	RedisClient *redis.Client
	DB          *sql.DB
}

func NewProductServiceServerImpl(db *sql.DB) *ProductServiceServerImpl {
	return &ProductServiceServerImpl{
		Repository:  repository.ProductRepositoryImpl{},
		RedisClient: setup.NewRedisClient(),
		DB:          db,
	}
}

func (server ProductServiceServerImpl) FindById(ctx context.Context, id *ProductId) (*ProductResponse, error) {
	tx, err := server.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	var productModel *model.ProductModel
	cache, err := server.RedisClient.Get(ctx, id.GetId()).Result()
	utils.LogIfError(err)

	if cache != "" {
		log.Printf("found in redis")
		err := json.Unmarshal([]byte(cache), &productModel)
		utils.LogIfError(err)
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
		utils.LogIfError(err)
		err = server.RedisClient.Set(ctx, id.GetId(), bytes, 1*time.Hour).Err()
		utils.LogIfError(err)
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
	tx, err := server.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

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
	tx, err := server.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

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
	tx, err := server.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

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
	tx, err := server.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

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
		utils.PanicIfError(err)

		if imageUrl.Valid {
			productTmp.ImageUrl = imageUrl.String
		}

		products = append(products, &productTmp)
	}
	utils.LogIfError(rows.Close())
	return products
}

func (server ProductServiceServerImpl) Save(ctx context.Context, product *Product) (*response.Response, error) {
	tx, err := server.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	categoryRepository := repository2.NewCategoryRepository()
	categoryModel := categoryRepository.FindById(ctx, tx, product.CategoryId)
	if categoryModel == nil {
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
	tx, err := server.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	categoryRepository := repository2.NewCategoryRepository()
	categoryModel := categoryRepository.FindById(ctx, tx, product.CategoryId)
	if categoryModel == nil {
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
	err = server.RedisClient.Del(ctx, product.Id).Err()
	utils.LogIfError(err)
	err = server.Repository.Update(ctx, tx, &productModel)
	status, message := setup.ResponseForModifying(err == nil)
	return &response.Response{
		Status:  status,
		Message: message,
	}, nil
}

func (server ProductServiceServerImpl) Delete(ctx context.Context, id *ProductId) (*response.Response, error) {
	tx, err := server.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)
	err = server.RedisClient.Del(ctx, id.GetId()).Err()
	utils.LogIfError(err)
	err = server.Repository.Delete(ctx, tx, id.Id)
	status, message := setup.ResponseForModifying(err == nil)
	return &response.Response{Status: status, Message: message}, nil
}

func (server ProductServiceServerImpl) mustEmbedUnimplementedProductServiceServer() {
	//TODO implement me
	panic("implement me")
}

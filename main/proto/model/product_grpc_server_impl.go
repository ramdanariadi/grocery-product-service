package model

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/ramdanariadi/grocery-be-golang/main/helpers"
	"github.com/ramdanariadi/grocery-be-golang/main/models"
	"github.com/ramdanariadi/grocery-be-golang/main/repositories/product"
	"github.com/ramdanariadi/grocery-be-golang/main/requestBody"
	"github.com/ramdanariadi/grocery-be-golang/main/utils"
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

func (server ProductServiceServerImpl) FindById(ctx context.Context, id *ProductId) (*ResponseWithData, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	var productModel models.ProductModel
	cache, _ := server.RedisClient.Get(ctx, id.GetId()).Result()
	helpers.LogIfError(err)

	if cache != "" {
		err := json.Unmarshal([]byte(cache), &productModel)
		helpers.LogIfError(err)
	} else {
		productModel = server.Repository.FindById(ctx, tx, id.Id)
		bytes, err := json.Marshal(productModel)
		helpers.LogIfError(err)
		err = server.RedisClient.Set(ctx, id.GetId(), bytes, 1*time.Hour).Err()
		helpers.LogIfError(err)
	}

	imageUrl := ""
	if str, ok := productModel.ImageUrl.(string); ok {
		imageUrl = str
	}
	grpcProductModel := &Product{
		Id:         productModel.Id,
		Name:       productModel.Name,
		Weight:     uint32(productModel.Weight),
		Category:   productModel.Category,
		ImageUrl:   imageUrl,
		CategoryId: productModel.CategoryId,
		Price:      uint64(productModel.Price),
	}
	return &ResponseWithData{
		Message: "OK",
		Status:  true,
		Data:    grpcProductModel,
	}, nil
}

func (server ProductServiceServerImpl) FindAll(ctx context.Context, _ *ProductEmpty) (*MultipleDataResponse, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	products := server.Repository.FindAll(ctx, tx)
	var data []*Product

	for _, productModel := range products {
		var imageUrl string
		if str, ok := productModel.ImageUrl.(string); ok {
			imageUrl = str
		}

		data = append(data, &Product{
			Id:         productModel.Id,
			Name:       productModel.Name,
			Price:      uint64(productModel.Price),
			CategoryId: productModel.CategoryId,
			Category:   productModel.Category,
			ImageUrl:   imageUrl,
			Weight:     uint32(productModel.Weight),
		})
	}

	return &MultipleDataResponse{
		Status:  len(products) > 0,
		Data:    data,
		Message: "",
	}, nil
}

func (server ProductServiceServerImpl) Save(ctx context.Context, product *Product) (*Response, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	newProduct := requestBody.ProductSaveRequest{
		Name:        product.Name,
		Price:       int64(product.Price),
		Weight:      uint(product.Weight),
		PerUnit:     int(product.PerUnit),
		Description: product.Description,
		Category:    product.CategoryId,
		ImageUrl:    product.ImageUrl,
	}
	saved := server.Repository.Save(ctx, tx, newProduct)
	return &Response{
		Status:  saved,
		Message: "OK",
	}, nil
}

func (server ProductServiceServerImpl) Update(ctx context.Context, product *Product) (*Response, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	updateProduct := requestBody.ProductSaveRequest{
		Name:        product.Name,
		Price:       int64(product.Price),
		Weight:      uint(product.Weight),
		PerUnit:     int(product.PerUnit),
		Description: product.Description,
	}
	updated := server.Repository.Update(ctx, tx, updateProduct, product.Id)
	return &Response{
		Status:  updated,
		Message: "OK",
	}, nil
}

func (server ProductServiceServerImpl) Delete(ctx context.Context, id *ProductId) (*Response, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	status := server.Repository.Delete(ctx, tx, id.Id)
	defer helpers.CommitOrRollback(tx)
	return &Response{Status: status, Message: "OK"}, nil
}

func (server ProductServiceServerImpl) mustEmbedUnimplementedProductServiceServer() {
	//TODO implement me
	panic("implement me")
}

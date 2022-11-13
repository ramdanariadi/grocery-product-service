package product

import (
	"database/sql"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/helpers"
	"github.com/ramdanariadi/grocery-product-service/main/models"
	"github.com/ramdanariadi/grocery-product-service/main/repositories/product"
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
	grpcProductModel := Product{
		Id:         productModel.Id,
		Name:       productModel.Name,
		Weight:     productModel.Weight,
		Category:   productModel.Category,
		ImageUrl:   imageUrl,
		CategoryId: productModel.CategoryId,
		Price:      productModel.Price,
	}
	return &ResponseWithData{
		Message: "OK",
		Status:  "Success",
		Data:    &grpcProductModel,
	}, nil
}

func (server ProductServiceServerImpl) FindProductsByCategory(ctx context.Context, id *CategoryId) (*MultipleDataResponse, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	if err != nil {
		return nil, err
	}
	productsRaw := server.Repository.FindByCategory(ctx, tx, id.Id)
	products := fetchProducts(productsRaw)
	status, message := utils.FetchResponseForCollection(len(products) > 0)
	return &MultipleDataResponse{
		Status:  status,
		Data:    products,
		Message: message,
	}, nil
}

func (server ProductServiceServerImpl) FindAll(ctx context.Context, _ *ProductEmpty) (*MultipleDataResponse, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	productsRaw := server.Repository.FindAll(ctx, tx)
	products := fetchProducts(productsRaw)
	status, message := utils.FetchResponseForCollection(len(products) > 0)
	return &MultipleDataResponse{
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

func (server ProductServiceServerImpl) Save(ctx context.Context, product *Product) (*Response, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	id, _ := uuid.NewUUID()
	productModel := models.ProductModel{
		Id:          id.String(),
		Name:        product.Name,
		Weight:      product.Weight,
		Category:    product.Category,
		CategoryId:  product.CategoryId,
		Price:       product.Price,
		Description: product.Description,
		ImageUrl:    product.ImageUrl,
	}
	saved := server.Repository.Save(ctx, tx, productModel)
	status, message := utils.FetchResponseForModifying(saved)
	return &Response{
		Status:  status,
		Message: message,
	}, nil
}

func (server ProductServiceServerImpl) Update(ctx context.Context, product *Product) (*Response, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	productModel := models.ProductModel{
		Id:          product.Id,
		Name:        product.Name,
		Weight:      product.Weight,
		Category:    product.Category,
		CategoryId:  product.CategoryId,
		Price:       product.Price,
		Description: product.Description,
		ImageUrl:    product.ImageUrl,
	}
	updated := server.Repository.Update(ctx, tx, productModel)
	status, message := utils.FetchResponseForModifying(updated)
	return &Response{
		Status:  status,
		Message: message,
	}, nil
}

func (server ProductServiceServerImpl) Delete(ctx context.Context, id *ProductId) (*Response, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	deleted := server.Repository.Delete(ctx, tx, id.Id)
	status, message := utils.FetchResponseForModifying(deleted)
	defer helpers.CommitOrRollback(tx)
	return &Response{Status: status, Message: message}, nil
}

func (server ProductServiceServerImpl) mustEmbedUnimplementedProductServiceServer() {
	//TODO implement me
	panic("implement me")
}

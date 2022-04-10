package model

import (
	"context"
	"database/sql"
	helpers2 "github.com/ramdanariadi/grocery-be-golang/main/helpers"
	productrepo "github.com/ramdanariadi/grocery-be-golang/main/repositories/product"
	"github.com/ramdanariadi/grocery-be-golang/main/requestBody"
)

type ProductServiceServerImpl struct {
	Repository productrepo.ProductRepositoryImpl
}

func NewProductServiceServerImpl(db *sql.DB) *ProductServiceServerImpl {
	return &ProductServiceServerImpl{
		Repository: productrepo.ProductRepositoryImpl{
			DB: db,
		},
	}
}

func (server ProductServiceServerImpl) FindById(ctx context.Context, id *ProductId) (*ResponseWithData, error) {
	tx, err := server.Repository.DB.Begin()
	helpers2.PanicIfError(err)
	defer helpers2.CommitOrRollback(tx)

	productModel := server.Repository.FindById(ctx, tx, id.Id)
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

func (server ProductServiceServerImpl) FindAll(ctx context.Context, empty *ProductEmpty) (*MultipleDataResponse, error) {
	tx, err := server.Repository.DB.Begin()
	helpers2.PanicIfError(err)
	defer helpers2.CommitOrRollback(tx)

	products := server.Repository.FindAll(ctx, tx)
	var data []*Product

	for _, product := range products {
		var imageUrl string
		if str, ok := product.ImageUrl.(string); ok {
			imageUrl = str
		}

		data = append(data, &Product{
			Id:         product.Id,
			Name:       product.Name,
			Price:      uint64(product.Price),
			CategoryId: product.CategoryId,
			Category:   product.Category,
			ImageUrl:   imageUrl,
			Weight:     uint32(product.Weight),
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
	helpers2.PanicIfError(err)
	defer helpers2.CommitOrRollback(tx)

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
	helpers2.PanicIfError(err)
	defer helpers2.CommitOrRollback(tx)

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
	helpers2.PanicIfError(err)
	status := server.Repository.Delete(ctx, tx, id.Id)
	defer helpers2.CommitOrRollback(tx)
	return &Response{Status: status, Message: "OK"}, nil
}

func (server ProductServiceServerImpl) mustEmbedUnimplementedProductServiceServer() {
	//TODO implement me
	panic("implement me")
}

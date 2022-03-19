package product

import (
	"context"
	"go-tunas/helpers"
	"go-tunas/models"
	"go-tunas/repositories/product"
	"go-tunas/requestBody"
)

type ProductServiceImpl struct {
	Repository product.ProductRepositoryImpl
}

func (service ProductServiceImpl) FindById(id string) models.ProductModel {
	tx, err := service.Repository.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	helpers.PanicIfError(err)
	return service.Repository.FindById(context.Background(), tx, id)
}

func (service ProductServiceImpl) FindAll() []models.ProductModel {
	tx, err := service.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	return service.Repository.FindAll(context.Background(), tx)
}

func (service ProductServiceImpl) Save(request requestBody.ProductSaveRequest) bool {
	tx, err := service.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	return service.Repository.Save(context.Background(), tx, request)
}

func (service ProductServiceImpl) Update(request requestBody.ProductSaveRequest, id string) bool {
	tx, err := service.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	return service.Repository.Update(context.Background(), tx, request, id)
}

func (service ProductServiceImpl) Delete(id string) bool {
	tx, err := service.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	return service.Repository.Delete(context.Background(), tx, id)
}

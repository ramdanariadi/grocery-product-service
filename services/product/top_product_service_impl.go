package product

import (
	"context"
	"go-tunas/helpers"
	"go-tunas/models"
	"go-tunas/repositories/product"
	"go-tunas/requestBody"
)

type TopProductServiceImpl struct {
	Repository product.TopProductRepositoryImpl
}

func (service TopProductServiceImpl) FindById(id string) models.ProductModel {
	tx, err := service.Repository.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	helpers.PanicIfError(err)
	return service.Repository.FindById(context.Background(), tx, id)
}

func (service TopProductServiceImpl) FindAll() []models.ProductModel {
	tx, err := service.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	return service.Repository.FindAll(context.Background(), tx)
}

func (service TopProductServiceImpl) Save(request requestBody.TopProductSaveRequest) bool {
	tx, err := service.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	return service.Repository.Save(context.Background(), tx, request)
}

func (service TopProductServiceImpl) Update(request requestBody.TopProductSaveRequest, id string) bool {
	tx, err := service.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	return service.Repository.Update(context.Background(), tx, request, id)
}

func (service TopProductServiceImpl) Delete(id string) bool {
	tx, err := service.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	return service.Repository.Delete(context.Background(), tx, id)
}

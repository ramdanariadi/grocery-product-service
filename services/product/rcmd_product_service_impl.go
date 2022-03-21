package product

import (
	"context"
	"go-tunas/helpers"
	"go-tunas/models"
	"go-tunas/repositories/product"
	"go-tunas/requestBody"
)

type RcmdProductServiceImpl struct {
	Repository product.RcmdProductRepositoryImpl
}

func (service RcmdProductServiceImpl) FindById(id string) models.ProductModel {
	tx, err := service.Repository.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	helpers.PanicIfError(err)
	return service.Repository.FindById(context.Background(), tx, id)
}

func (service RcmdProductServiceImpl) FindAll() []models.ProductModel {
	tx, err := service.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	return service.Repository.FindAll(context.Background(), tx)
}

func (service RcmdProductServiceImpl) Save(request requestBody.RcmdProductSaveRequest) bool {
	tx, err := service.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	return service.Repository.Save(context.Background(), tx, request)
}

func (service RcmdProductServiceImpl) Update(request requestBody.RcmdProductSaveRequest, id string) bool {
	tx, err := service.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	return service.Repository.Update(context.Background(), tx, request, id)
}

func (service RcmdProductServiceImpl) Delete(id string) bool {
	tx, err := service.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	return service.Repository.Delete(context.Background(), tx, id)
}

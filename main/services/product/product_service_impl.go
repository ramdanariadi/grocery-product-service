package product

import (
	"context"
	helpers2 "github.com/ramdanariadi/grocery-be-golang/main/helpers"
	"github.com/ramdanariadi/grocery-be-golang/main/models"
	"github.com/ramdanariadi/grocery-be-golang/main/repositories/product"
	"github.com/ramdanariadi/grocery-be-golang/main/requestBody"
)

type ProductServiceImpl struct {
	Repository product.ProductRepositoryImpl
}

func (service ProductServiceImpl) FindById(id string) models.ProductModel {
	tx, err := service.Repository.DB.Begin()
	defer helpers2.CommitOrRollback(tx)
	helpers2.PanicIfError(err)
	return service.Repository.FindById(context.Background(), tx, id)
}

func (service ProductServiceImpl) FindAll() []models.ProductModel {
	tx, err := service.Repository.DB.Begin()
	helpers2.PanicIfError(err)
	defer helpers2.CommitOrRollback(tx)
	return service.Repository.FindAll(context.Background(), tx)
}

func (service ProductServiceImpl) Save(request requestBody.ProductSaveRequest) bool {
	tx, err := service.Repository.DB.Begin()
	helpers2.PanicIfError(err)
	defer helpers2.CommitOrRollback(tx)
	return service.Repository.Save(context.Background(), tx, request)
}

func (service ProductServiceImpl) Update(request requestBody.ProductSaveRequest, id string) bool {
	tx, err := service.Repository.DB.Begin()
	helpers2.PanicIfError(err)
	defer helpers2.CommitOrRollback(tx)
	return service.Repository.Update(context.Background(), tx, request, id)
}

func (service ProductServiceImpl) Delete(id string) bool {
	tx, err := service.Repository.DB.Begin()
	helpers2.PanicIfError(err)
	defer helpers2.CommitOrRollback(tx)
	return service.Repository.Delete(context.Background(), tx, id)
}

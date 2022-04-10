package category

import (
	"context"
	helpers2 "github.com/ramdanariadi/grocery-be-golang/main/helpers"
	categoryModel "github.com/ramdanariadi/grocery-be-golang/main/models"
	"github.com/ramdanariadi/grocery-be-golang/main/repositories/category"
	"github.com/ramdanariadi/grocery-be-golang/main/requestBody"
)

type CategoryServiceImpl struct {
	Repository category.CategoryRepositoryImpl
}

func (service CategoryServiceImpl) FindById(id string) categoryModel.CategoryModel {
	bgctx := context.Background()
	tx, err := service.Repository.DB.Begin()
	helpers2.PanicIfError(err)
	return service.Repository.FindById(bgctx, tx, id)
}

func (service CategoryServiceImpl) FindAll() []categoryModel.CategoryModel {
	bgctx := context.Background()
	tx, err := service.Repository.DB.Begin()
	helpers2.PanicIfError(err)
	return service.Repository.FindAll(bgctx, tx)
}

func (service CategoryServiceImpl) Save(request requestBody.CategorySaveRequest) bool {
	bgctx := context.Background()
	tx, err := service.Repository.DB.Begin()
	helpers2.PanicIfError(err)
	defer helpers2.CommitOrRollback(tx)
	return service.Repository.Save(bgctx, tx, request)
}

func (service CategoryServiceImpl) Update(request requestBody.CategorySaveRequest, id string) bool {
	bgcontext := context.Background()
	tx, err := service.Repository.DB.Begin()
	helpers2.PanicIfError(err)
	defer helpers2.CommitOrRollback(tx)

	return service.Repository.Update(bgcontext, tx, request, id)
}

func (service CategoryServiceImpl) Delete(id string) bool {
	bgcontext := context.Background()
	tx, err := service.Repository.DB.Begin()
	helpers2.PanicIfError(err)
	defer helpers2.CommitOrRollback(tx)
	return service.Repository.Delete(bgcontext, tx, id)
}

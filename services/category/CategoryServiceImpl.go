package category

import (
	"context"
	categoryModel "go-tunas/models/category"
	"go-tunas/repositories/category"
	"go-tunas/requestBody"
)

type CategoryServiceImpl struct {
	Repository category.CategoryRepositoryImpl
}

func (service CategoryServiceImpl) FindById(id string) categoryModel.CategoryModel {
	//TODO implement me
	panic("implement me")
}

func (service CategoryServiceImpl) FindAll() []categoryModel.CategoryModel {
	bgctx := context.Background()
	tx, err := service.Repository.DB.Begin()

	if err != nil {
		panic("tx error")
	}

	return service.Repository.FindAll(bgctx, tx)
}

func (service CategoryServiceImpl) Save(request requestBody.CategorySaveRequest) bool {
	//TODO implement me
	panic("implement me")
}

func (service CategoryServiceImpl) Update(request requestBody.CategorySaveRequest, id string) bool {
	//TODO implement me
	panic("implement me")
}

func (service CategoryServiceImpl) Delete(id string) bool {
	//TODO implement me
	panic("implement me")
}

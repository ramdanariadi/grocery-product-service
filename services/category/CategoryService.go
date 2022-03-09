package category

import (
	"go-tunas/models/category"
	"go-tunas/requestBody"
)

type CategoryService interface {
	FindById(id string) category.CategoryModel
	FindAll() []category.CategoryModel
	Save(request requestBody.CategorySaveRequest) bool
	Update(request requestBody.CategorySaveRequest, id string) bool
	Delete(id string) bool
}

package category

import (
	"go-tunas/models"
	"go-tunas/requestBody"
)

type CategoryService interface {
	FindById(id string) models.CategoryModel
	FindAll() []models.CategoryModel
	Save(request requestBody.CategorySaveRequest) bool
	Update(request requestBody.CategorySaveRequest, id string) bool
	Delete(id string) bool
}

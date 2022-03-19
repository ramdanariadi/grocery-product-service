package product

import (
	"go-tunas/models"
	"go-tunas/requestBody"
)

type ProductService interface {
	FindById(id string) models.ProductModel
	FindAll() []models.ProductModel
	Save(request requestBody.ProductSaveRequest) bool
	Update(request requestBody.ProductSaveRequest, id string) bool
	Delete(id string) bool
}

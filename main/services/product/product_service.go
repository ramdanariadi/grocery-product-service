package product

import (
	"github.com/ramdanariadi/grocery-be-golang/main/models"
	"github.com/ramdanariadi/grocery-be-golang/main/requestBody"
)

type ProductService interface {
	FindById(id string) models.ProductModel
	FindAll() []models.ProductModel
	Save(request requestBody.ProductSaveRequest) bool
	Update(request requestBody.ProductSaveRequest, id string) bool
	Delete(id string) bool
}

type TopProductService interface {
	FindById(id string) models.ProductModel
	FindAll() []models.ProductModel
	Save(id string) bool
	Delete(id string) bool
}

type RcmdProductService interface {
	FindById(id string) models.ProductModel
	FindAll() []models.ProductModel
	Save(id string) bool
	Delete(id string) bool
}

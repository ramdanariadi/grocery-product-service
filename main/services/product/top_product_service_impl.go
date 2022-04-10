package product

import (
	"context"
	helpers2 "github.com/ramdanariadi/grocery-be-golang/main/helpers"
	"github.com/ramdanariadi/grocery-be-golang/main/models"
	product2 "github.com/ramdanariadi/grocery-be-golang/main/repositories/product"
	"github.com/ramdanariadi/grocery-be-golang/main/requestBody"
)

type TopProductServiceImpl struct {
	TopProductRepository product2.TopProductRepositoryImpl
	ProductRepository    product2.ProductRepositoryImpl
}

func (service TopProductServiceImpl) FindById(id string) models.ProductModel {
	tx, err := service.TopProductRepository.DB.Begin()
	defer helpers2.CommitOrRollback(tx)
	helpers2.PanicIfError(err)
	return service.TopProductRepository.FindById(context.Background(), tx, id)
}

func (service TopProductServiceImpl) FindAll() []models.ProductModel {
	tx, err := service.TopProductRepository.DB.Begin()
	helpers2.PanicIfError(err)
	defer helpers2.CommitOrRollback(tx)
	return service.TopProductRepository.FindAll(context.Background(), tx)
}

func (service TopProductServiceImpl) Save(id string) bool {
	tx, err := service.TopProductRepository.DB.Begin()
	helpers2.PanicIfError(err)
	defer helpers2.CommitOrRollback(tx)

	bgContext := context.Background()
	productModel := service.ProductRepository.FindById(bgContext, tx, id)

	if productModel == (models.ProductModel{}) {
		return false
	}

	topProduct := requestBody.TopProductSaveRequest{
		ProductId:   productModel.Id,
		Name:        productModel.Name,
		Weight:      productModel.Weight,
		Price:       productModel.Price,
		PerUnit:     productModel.PerUnit,
		Category:    productModel.Category,
		ImageUrl:    productModel.ImageUrl,
		Description: productModel.Description,
	}

	return service.TopProductRepository.Save(context.Background(), tx, topProduct)
}

func (service TopProductServiceImpl) Delete(id string) bool {
	tx, err := service.TopProductRepository.DB.Begin()
	helpers2.PanicIfError(err)
	defer helpers2.CommitOrRollback(tx)
	return service.TopProductRepository.Delete(context.Background(), tx, id)
}

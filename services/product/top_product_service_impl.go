package product

import (
	"context"
	"go-tunas/helpers"
	"go-tunas/models"
	"go-tunas/repositories/product"
	"go-tunas/requestBody"
)

type TopProductServiceImpl struct {
	TopProductRepository product.TopProductRepositoryImpl
	ProductRepository    product.ProductRepositoryImpl
}

func (service TopProductServiceImpl) FindById(id string) models.ProductModel {
	tx, err := service.TopProductRepository.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	helpers.PanicIfError(err)
	return service.TopProductRepository.FindById(context.Background(), tx, id)
}

func (service TopProductServiceImpl) FindAll() []models.ProductModel {
	tx, err := service.TopProductRepository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	return service.TopProductRepository.FindAll(context.Background(), tx)
}

func (service TopProductServiceImpl) Save(id string) bool {
	tx, err := service.TopProductRepository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

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
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	return service.TopProductRepository.Delete(context.Background(), tx, id)
}

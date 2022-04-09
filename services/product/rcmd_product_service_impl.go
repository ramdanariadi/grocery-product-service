package product

import (
	"context"
	"github.com/ramdanariadi/grocery-be-golang/helpers"
	"github.com/ramdanariadi/grocery-be-golang/models"
	"github.com/ramdanariadi/grocery-be-golang/repositories/product"
	"github.com/ramdanariadi/grocery-be-golang/requestBody"
)

type RcmdProductServiceImpl struct {
	RcmdRepository    product.RcmdProductRepositoryImpl
	ProductRepository product.ProductRepositoryImpl
}

func (service RcmdProductServiceImpl) FindById(id string) models.ProductModel {
	tx, err := service.RcmdRepository.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	helpers.PanicIfError(err)
	return service.RcmdRepository.FindById(context.Background(), tx, id)
}

func (service RcmdProductServiceImpl) FindAll() []models.ProductModel {
	tx, err := service.RcmdRepository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	return service.RcmdRepository.FindAll(context.Background(), tx)
}

func (service RcmdProductServiceImpl) Save(id string) bool {
	tx, err := service.RcmdRepository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	bgContext := context.Background()

	product := service.ProductRepository.FindById(bgContext, tx, id)
	productSaveRequest := requestBody.RcmdProductSaveRequest{
		ProductId:   product.Id,
		Price:       product.Price,
		Description: product.Description,
		ImageUrl:    product.ImageUrl,
		Category:    product.Category,
		PerUnit:     product.PerUnit,
		Weight:      product.Weight,
		Name:        product.Name,
	}

	return service.RcmdRepository.Save(context.Background(), tx, productSaveRequest)
}

func (service RcmdProductServiceImpl) Delete(id string) bool {
	tx, err := service.RcmdRepository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	return service.RcmdRepository.Delete(context.Background(), tx, id)
}

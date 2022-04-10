package product

import (
	"context"
	helpers2 "github.com/ramdanariadi/grocery-be-golang/main/helpers"
	"github.com/ramdanariadi/grocery-be-golang/main/models"
	product2 "github.com/ramdanariadi/grocery-be-golang/main/repositories/product"
	"github.com/ramdanariadi/grocery-be-golang/main/requestBody"
)

type RcmdProductServiceImpl struct {
	RcmdRepository    product2.RcmdProductRepositoryImpl
	ProductRepository product2.ProductRepositoryImpl
}

func (service RcmdProductServiceImpl) FindById(id string) models.ProductModel {
	tx, err := service.RcmdRepository.DB.Begin()
	defer helpers2.CommitOrRollback(tx)
	helpers2.PanicIfError(err)
	return service.RcmdRepository.FindById(context.Background(), tx, id)
}

func (service RcmdProductServiceImpl) FindAll() []models.ProductModel {
	tx, err := service.RcmdRepository.DB.Begin()
	helpers2.PanicIfError(err)
	defer helpers2.CommitOrRollback(tx)
	return service.RcmdRepository.FindAll(context.Background(), tx)
}

func (service RcmdProductServiceImpl) Save(id string) bool {
	tx, err := service.RcmdRepository.DB.Begin()
	helpers2.PanicIfError(err)
	defer helpers2.CommitOrRollback(tx)

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
	helpers2.PanicIfError(err)
	defer helpers2.CommitOrRollback(tx)
	return service.RcmdRepository.Delete(context.Background(), tx, id)
}

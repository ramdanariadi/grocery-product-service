package product

import (
	"context"
	"github.com/google/uuid"
	productresponse "go-tunas/customresponses/product"
	"go-tunas/helpers"
	"go-tunas/models"
	"go-tunas/repositories/product"
	"go-tunas/repositories/transactions"
)

type WishlistServiceImpl struct {
	WishlistRepository transactions.WishlistRepositoryImpl
	ProductRepository  product.ProductRepositoryImpl
}

func (service WishlistServiceImpl) FindByUserId(userId string) []productresponse.WishlistResponse {
	tx, err := service.WishlistRepository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	return service.WishlistRepository.FindByUserId(context.Background(), tx, userId)
}

func (service WishlistServiceImpl) FindByUserAndProductId(userId string, productId string) productresponse.WishlistResponse {
	tx, err := service.WishlistRepository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	return service.WishlistRepository.FindByUserAndProductId(context.Background(), tx, userId, productId)
}

func (service WishlistServiceImpl) Save(userId string, productId string) bool {
	tx, err := service.WishlistRepository.DB.Begin()
	ctx := context.Background()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	product := service.ProductRepository.FindById(ctx, tx, productId)
	id, err := uuid.NewUUID()
	helpers.PanicIfError(err)
	wishlistModel := models.WishlistModel{
		Id:        id.String(),
		Name:      product.Name,
		Weight:    product.Weight,
		Price:     product.Price,
		PerUnit:   product.PerUnit,
		Category:  product.Category,
		ImageUrl:  product.ImageUrl,
		ProductId: product.Id,
		UserId:    userId,
	}
	return service.WishlistRepository.Save(ctx, tx, wishlistModel)
}

func (service WishlistServiceImpl) Delete(userId string, productId string) bool {
	tx, err := service.WishlistRepository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	return service.WishlistRepository.Delete(context.Background(), tx, userId, productId)
}

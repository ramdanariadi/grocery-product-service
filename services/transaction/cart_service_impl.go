package transaction

import (
	"context"
	"github.com/google/uuid"
	"go-tunas/customresponses/product"
	"go-tunas/helpers"
	"go-tunas/models"
	productrepo "go-tunas/repositories/product"
	"go-tunas/repositories/transactions"
)

type CartServiceImpl struct {
	CartRepository    transactions.CartRepositoryImpl
	ProductRepository productrepo.ProductRepositoryImpl
}

func (service CartServiceImpl) FindById(id string) []product.CartResponse {
	tx, err := service.CartRepository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	carts := service.CartRepository.FindByUserId(context.Background(), tx, id)
	var cartsResponse []product.CartResponse
	for _, cart := range carts {
		product := product.CartResponse{
			Id:       cart.Id,
			Name:     cart.Name,
			Weight:   cart.Weight,
			Price:    cart.Price,
			PerUnit:  cart.PerUnit,
			Category: cart.Category,
			ImageUrl: cart.ImageUrl,
			Total:    cart.Total,
		}
		cartsResponse = append(cartsResponse, product)
	}
	return cartsResponse
}

func (service CartServiceImpl) Save(userId string, productId string, total int) bool {
	tx, err := service.CartRepository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	ctx := context.Background()
	product := service.ProductRepository.FindById(ctx, tx, productId)
	id, err := uuid.NewUUID()
	helpers.PanicIfError(err)
	cartModel := models.CartModel{
		Id:        id.String(),
		Name:      product.Name,
		Weight:    product.Weight,
		Price:     product.Price,
		PerUnit:   product.PerUnit,
		Category:  product.Category,
		ImageUrl:  product.ImageUrl,
		UserId:    userId,
		ProductId: product.Id,
		Total:     total,
	}
	return service.CartRepository.Save(ctx, tx, cartModel)
}

func (service CartServiceImpl) Delete(userId string, produtId string) bool {
	tx, err := service.CartRepository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	return service.CartRepository.Delete(context.Background(), tx, userId, produtId)
}

package transaction

import (
	"context"
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-be-golang/main/customresponses/product"
	helpers2 "github.com/ramdanariadi/grocery-be-golang/main/helpers"
	"github.com/ramdanariadi/grocery-be-golang/main/models"
	productrepo "github.com/ramdanariadi/grocery-be-golang/main/repositories/product"
	"github.com/ramdanariadi/grocery-be-golang/main/repositories/transactions"
)

type CartServiceImpl struct {
	CartRepository    transactions.CartRepositoryImpl
	ProductRepository productrepo.ProductRepositoryImpl
}

func (service CartServiceImpl) FindById(id string) []product.CartResponse {
	tx, err := service.CartRepository.DB.Begin()
	helpers2.PanicIfError(err)
	defer helpers2.CommitOrRollback(tx)
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
	helpers2.PanicIfError(err)
	defer helpers2.CommitOrRollback(tx)
	ctx := context.Background()
	product := service.ProductRepository.FindById(ctx, tx, productId)
	id, err := uuid.NewUUID()
	helpers2.PanicIfError(err)
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
	helpers2.PanicIfError(err)
	defer helpers2.CommitOrRollback(tx)
	return service.CartRepository.Delete(context.Background(), tx, userId, produtId)
}

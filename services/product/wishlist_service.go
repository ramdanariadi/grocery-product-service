package product

import (
	"go-tunas/customresponses/product"
)

type WishlistService interface {
	FindByUserId(userId string) []product.WishlistResponse
	FindByUserAndProductId(userId string, productId string) product.WishlistResponse
	Save(userId string, productId string) bool
	Delete(userId string, productId string) bool
}

package product

import (
	"github.com/ramdanariadi/grocery-be-golang/main/customresponses/product"
)

type WishlistService interface {
	FindByUserId(userId string) []product.WishlistResponse
	FindByUserAndProductId(userId string, productId string) product.WishlistResponse
	Save(userId string, productId string) bool
	Delete(userId string, productId string) bool
}

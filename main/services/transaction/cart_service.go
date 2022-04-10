package transaction

import (
	"github.com/ramdanariadi/grocery-be-golang/main/customresponses/product"
)

type CartService interface {
	FindById(id string) []product.CartResponse
	Save(userId string, produtId string, total int) bool
	Delete(userId string, produtId string) bool
}

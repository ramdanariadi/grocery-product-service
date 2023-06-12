package cart

import "github.com/ramdanariadi/grocery-product-service/main/cart/dto"

type Service interface {
	Store(productId string, total uint, userId string) *dto.CartTotalItemDTO
	Destroy(id string, userId string)
	Find(reqBody *dto.FindCartDTO) []*dto.Cart
}

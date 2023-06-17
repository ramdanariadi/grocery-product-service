package wishlist

import "github.com/ramdanariadi/grocery-product-service/main/wishlist/dto"

type Service interface {
	Store(productId string, userId string)
	Destroy(productId string, userId string)
	Find(reqBody *dto.FindWishlistDTO) []*dto.WishlistDTO
	FindByProductId(productId string, userId string) *dto.WishlistDTO
}

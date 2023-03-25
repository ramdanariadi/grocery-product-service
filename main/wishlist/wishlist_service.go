package wishlist

import "github.com/ramdanariadi/grocery-product-service/main/wishlist/dto"

type Service interface {
	Store(productId string, userId string)
	Destroy(id string, userId string)
	Find(reqBody *dto.FindWishlistDTO) []*dto.WishlistDTO
}

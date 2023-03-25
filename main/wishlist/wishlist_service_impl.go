package wishlist

import (
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/exception"
	"github.com/ramdanariadi/grocery-product-service/main/product"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"github.com/ramdanariadi/grocery-product-service/main/wishlist/dto"
	"gorm.io/gorm"
	"strings"
)

type ServiceImpl struct {
	DB *gorm.DB
}

func (service ServiceImpl) Store(productId string, userId string) {
	var p product.Product
	tx := service.DB.Where("id = ?", productId).Find(&p)
	if tx.Error != nil {
		panic(exception.ValidationException{Message: "INVALID_PRODUCT"})
	}
	id, _ := uuid.NewUUID()
	wishlist := Wishlist{
		ID:      id.String(),
		Product: p,
		UserId:  userId,
	}
	save := service.DB.Save(&wishlist)
	utils.PanicIfError(save.Error)
}

func (service ServiceImpl) Destroy(id string, userId string) {
	wishlist := Wishlist{ID: id}
	find := service.DB.Find(&wishlist)
	if find.Error != nil {
		panic(exception.ValidationException{"INVALID_WISHLIST"})
	}
	tx := service.DB.Delete(&wishlist)
	utils.PanicIfError(tx.Error)
}

func (service ServiceImpl) Find(reqBody *dto.FindWishlistDTO) []*dto.WishlistDTO {
	var wishlists []*Wishlist
	tx := service.DB.Model(&Wishlist{})
	tx.Joins("LEFT JOIN products p ON p.id = wishlists.product_id AND p.deleted_at IS NULL")
	tx.Joins("LEFT JOIN categories c ON c.id = p.category_id AND c.deleted_at IS NULL")
	tx.Preload("Product.Category")
	if reqBody.Search != nil {
		tx.Where("LOWER(p.name) like ?", strings.ToLower("%"+*reqBody.Search+"%"))
	}
	tx.Limit(reqBody.PageSize).Offset(reqBody.PageIndex * reqBody.PageSize).Find(&wishlists)
	wishlistsResult := make([]*dto.WishlistDTO, 0)
	for _, wishlist := range wishlists {
		wishlistsResult = append(wishlistsResult, &dto.WishlistDTO{
			ID:          wishlist.ID,
			Name:        wishlist.Product.Name,
			Category:    wishlist.Product.Category.Category,
			ImageUrl:    wishlist.Product.ImageUrl,
			Price:       wishlist.Product.Price,
			PerUnit:     wishlist.Product.PerUnit,
			Weight:      wishlist.Product.Weight,
			Description: wishlist.Product.Description,
		})
	}
	return wishlistsResult
}

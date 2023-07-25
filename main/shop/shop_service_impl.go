package shop

import (
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/exception"
	userModel "github.com/ramdanariadi/grocery-product-service/main/user"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"gorm.io/gorm"
	"log"
)

type ShopServiceImpl struct {
	*gorm.DB
}

func (service *ShopServiceImpl) AddShop(userId string, shop AddShopDTO) {
	if nil == shop.Address || nil == shop.Name {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}

	var user userModel.User
	find := service.DB.Where("id = ?", userId).Find(&user)
	if nil != find.Error {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}

	newUUID, _ := uuid.NewUUID()
	s := Shop{
		ID:       newUUID.String(),
		Name:     *shop.Name,
		Address:  *shop.Address,
		User:     user,
		ImageUrl: *shop.ImageUrl,
	}
	tx := service.DB.Create(&s)
	utils.PanicIfError(tx.Error)
}

func (service *ShopServiceImpl) UpdateShop(userId string, reqBody EditShopDTO) {
	var shop Shop
	tx := service.DB.Where("user_id = ? ", userId).Find(&shop)
	if tx.RowsAffected < 1 {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}

	shop.Name = *reqBody.Name
	shop.Address = *reqBody.Address
	shop.ImageUrl = *reqBody.ImageUrl

	save := service.DB.Save(&shop)
	utils.PanicIfError(save.Error)
}

func (service *ShopServiceImpl) GetShop(userId string) ShopDTO {
	var shop Shop
	log.Print("user id : " + userId)
	tx := service.DB.Where("user_id = ?", userId).Find(&shop)
	if tx.RowsAffected < 1 {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}

	return ShopDTO{Id: shop.ID, Name: shop.Name, Address: shop.Address, ImageUrl: shop.ImageUrl}
}

func (service *ShopServiceImpl) DeleteShop(userID string) {
	var shop Shop
	tx := service.DB.Where("user_id = ?", userID).Find(&shop)
	if tx.RowsAffected < 1 {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}

	service.DB.Delete(&shop)
}

package shop

import (
	"github.com/gin-gonic/gin"
	"github.com/ramdanariadi/grocery-product-service/main/exception"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"gorm.io/gorm"
)

type ShopControllerImpl struct {
	shopService ShopService
}

func NewShopController(db *gorm.DB) ShopController {
	return &ShopControllerImpl{shopService: &ShopServiceImpl{db}}
}

func (controller *ShopControllerImpl) AddShop(ctx *gin.Context) {
	var requestBody AddShopDTO
	err := ctx.ShouldBind(&requestBody)
	utils.PanicIfError(err)

	value, exists := ctx.Get("userId")
	if !exists {
		panic(exception.AuthenticationException{Message: exception.Unauthorized})
	}
	controller.shopService.AddShop(value.(string), requestBody)
	ctx.JSON(200, gin.H{})
}

func (controller *ShopControllerImpl) EditShop(ctx *gin.Context) {
	var requestBody EditShopDTO
	err := ctx.ShouldBind(&requestBody)
	utils.PanicIfError(err)

	value, exists := ctx.Get("userId")
	if !exists {
		panic(exception.AuthenticationException{Message: exception.Unauthorized})
	}
	controller.shopService.UpdateShop(value.(string), requestBody)
	ctx.JSON(200, gin.H{})
}

func (controller *ShopControllerImpl) DeleteShop(ctx *gin.Context) {
	value, exists := ctx.Get("userId")
	if !exists {
		panic(exception.AuthenticationException{Message: exception.Unauthorized})
	}
	controller.shopService.DeleteShop(value.(string))
	ctx.JSON(200, gin.H{})
}

func (controller *ShopControllerImpl) GetShop(ctx *gin.Context) {
	value, exists := ctx.Get("userId")
	if !exists {
		panic(exception.AuthenticationException{Message: exception.Unauthorized})
	}
	shop := controller.shopService.GetShop(value.(string))
	ctx.JSON(200, gin.H{"data": shop})
}

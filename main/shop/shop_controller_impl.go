package shop

import (
	"github.com/gin-gonic/gin"
	"github.com/ramdanariadi/grocery-product-service/main/exception"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"gorm.io/gorm"
)

type ControllerImpl struct {
	shopService ShopService
}

func NewShopController(db *gorm.DB) Controller {
	return &ControllerImpl{shopService: &ShopServiceImpl{db}}
}

func (controller *ControllerImpl) AddShop(ctx *gin.Context) {
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

func (controller *ControllerImpl) EditShop(ctx *gin.Context) {
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

func (controller *ControllerImpl) DeleteShop(ctx *gin.Context) {
	value, exists := ctx.Get("userId")
	if !exists {
		panic(exception.AuthenticationException{Message: exception.Unauthorized})
	}
	controller.shopService.DeleteShop(value.(string))
	ctx.JSON(200, gin.H{})
}

func (controller *ControllerImpl) GetShop(ctx *gin.Context) {
	value, exists := ctx.Get("userId")
	if !exists {
		panic(exception.AuthenticationException{Message: exception.Unauthorized})
	}
	shop := controller.shopService.GetShop(value.(string))
	ctx.JSON(200, gin.H{"data": shop})
}

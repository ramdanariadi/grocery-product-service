package wishlist

import (
	"github.com/gin-gonic/gin"
	"github.com/ramdanariadi/grocery-product-service/main/exception"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"github.com/ramdanariadi/grocery-product-service/main/wishlist/dto"
	"gorm.io/gorm"
)

type ControllerImpl struct {
	Service Service
}

func NewWishlistController(DB *gorm.DB) *ControllerImpl {
	return &ControllerImpl{Service: ServiceImpl{DB: DB}}
}

func (controller ControllerImpl) Store(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists {
		panic(exception.AuthenticationException{Message: "UNAUTHORIZED"})
	}
	productId := ctx.Param("productId")
	controller.Service.Store(productId, userId.(string))
	ctx.JSON(200, gin.H{})
}

func (controller ControllerImpl) Destroy(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists {
		panic(exception.AuthenticationException{Message: "UNAUTHORIZED"})
	}
	productId := ctx.Param("productId")
	controller.Service.Destroy(productId, userId.(string))
	ctx.JSON(200, gin.H{})
}

func (controller ControllerImpl) Find(ctx *gin.Context) {
	var reqBody dto.FindWishlistDTO
	err := ctx.ShouldBind(&reqBody)
	utils.PanicIfError(err)
	wishlists := controller.Service.Find(&reqBody)
	ctx.JSON(200, gin.H{"data": wishlists})
}

func (controller ControllerImpl) FindByProductId(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists {
		panic(exception.AuthenticationException{Message: "UNAUTHORIZED"})
	}
	productId := ctx.Param("productId")
	wishlist := controller.Service.FindByProductId(productId, userId.(string))
	ctx.JSON(200, gin.H{"data": wishlist})
}

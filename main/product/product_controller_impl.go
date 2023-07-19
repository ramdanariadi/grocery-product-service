package product

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/ramdanariadi/grocery-product-service/main/exception"
	"github.com/ramdanariadi/grocery-product-service/main/product/dto"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"gorm.io/gorm"
)

type ProductControllerImpl struct {
	Service ProductService
}

func NewProductController(db *gorm.DB, redisClient *redis.Client) *ProductControllerImpl {
	return &ProductControllerImpl{Service: ProductServiceImpl{DB: db, Redish: redisClient}}
}

func (controller ProductControllerImpl) Save(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists {
		panic(exception.AuthenticationException{Message: exception.Unauthorized})
	}
	var request dto.AddProductDTO
	err := ctx.Bind(&request)
	utils.LogIfError(err)
	controller.Service.Save(userId.(string), &request)
	ctx.JSON(200, gin.H{})
}

func (controller ProductControllerImpl) FindById(ctx *gin.Context) {
	id := ctx.Param("id")
	productDTO := controller.Service.FindById(id)
	ctx.JSON(200, gin.H{"data": productDTO})
}

func (controller ProductControllerImpl) FindAll(ctx *gin.Context) {
	var request dto.FindProductRequest
	err := ctx.ShouldBindQuery(&request)
	utils.PanicIfError(err)
	response := controller.Service.FindAll(&request)
	ctx.JSON(200, response)
}

func (controller ProductControllerImpl) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var request dto.AddProductDTO
	ctx.Bind(&request)
	controller.Service.Update(id, &request)
	ctx.JSON(200, gin.H{})
}

func (controller ProductControllerImpl) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	controller.Service.Delete(id)
	ctx.JSON(200, gin.H{})
}

func (controller ProductControllerImpl) SetTopProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	controller.Service.SetTop(id)
	ctx.JSON(200, gin.H{})
}

func (controller ProductControllerImpl) SetRecommendationProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	controller.Service.SetRecommendation(id)
	ctx.JSON(200, gin.H{})
}

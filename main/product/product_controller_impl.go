package product

import (
	"github.com/gin-gonic/gin"
	"github.com/ramdanariadi/grocery-product-service/main/product/dto"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"gorm.io/gorm"
)

type ProductControllerImpl struct {
	Service ProductService
}

func NewProductController(db *gorm.DB) *ProductControllerImpl {
	return &ProductControllerImpl{Service: ProductServiceImpl{DB: db}}
}

func (controller ProductControllerImpl) Save(ctx *gin.Context) {
	var request dto.AddProductDTO
	err := ctx.Bind(&request)
	utils.LogIfError(err)
	controller.Service.Save(&request)
	ctx.JSON(200, gin.H{})
}

func (controller ProductControllerImpl) FindById(ctx *gin.Context) {
	id := ctx.Param("id")
	productDTO := controller.Service.FindById(id)
	ctx.JSON(200, gin.H{"data": productDTO})
}

func (controller ProductControllerImpl) FindAll(ctx *gin.Context) {
	var request dto.FindProductRequest
	response := controller.Service.FindAll(&request)
	ctx.JSON(200, gin.H{"data": response})
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

}

func (controller ProductControllerImpl) TopProduct(ctx *gin.Context) {

}

func (controller ProductControllerImpl) SetRecommendationProduct(ctx *gin.Context) {

}

func (controller ProductControllerImpl) RecommendationProduct(ctx *gin.Context) {

}

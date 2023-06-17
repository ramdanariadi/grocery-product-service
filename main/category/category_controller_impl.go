package category

import (
	"github.com/gin-gonic/gin"
	"github.com/ramdanariadi/grocery-product-service/main/category/dto"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"gorm.io/gorm"
)

type CategoryControllerImpl struct {
	Service CategoryService
}

func NewCategoryController(db *gorm.DB) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		Service: CategoryServiceImpl{
			DB: db,
		},
	}
}

func (controller CategoryControllerImpl) FindAll(ctx *gin.Context) {
	var param dto.PaginationDTO
	err := ctx.ShouldBindQuery(&param)
	utils.PanicIfError(err)
	ctx.JSON(200, controller.Service.FindAll(param.PageIndex, param.PageSize))
}

func (controller CategoryControllerImpl) FindById(ctx *gin.Context) {
	id := ctx.Param("id")
	category := controller.Service.FindById(id)
	ctx.JSON(200, gin.H{"data": category})
}

func (controller CategoryControllerImpl) Save(ctx *gin.Context) {
	request := dto.AddCategoryDTO{}
	err := ctx.Bind(&request)
	utils.PanicIfError(err)
	controller.Service.Save(&request)
	ctx.JSON(200, gin.H{})
}

func (controller CategoryControllerImpl) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var request dto.AddCategoryDTO
	ctx.Bind(&request)
	controller.Service.Update(id, &request)
	ctx.JSON(200, gin.H{})
}

func (controller CategoryControllerImpl) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	controller.Service.Delete(id)
	ctx.JSON(200, gin.H{})
}

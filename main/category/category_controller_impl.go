package category

import (
	"github.com/gin-gonic/gin"
	"github.com/ramdanariadi/grocery-product-service/main/category/dto"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"gorm.io/gorm"
)

type ControllerImpl struct {
	Service CategoryService
}

func NewCategoryController(db *gorm.DB) *ControllerImpl {
	return &ControllerImpl{
		Service: CategoryServiceImpl{
			DB: db,
		},
	}
}

func (controller ControllerImpl) FindAll(ctx *gin.Context) {
	var param dto.PaginationDTO
	err := ctx.ShouldBindQuery(&param)
	utils.PanicIfError(err)
	ctx.JSON(200, controller.Service.FindAll(param.PageIndex, param.PageSize))
}

func (controller ControllerImpl) FindById(ctx *gin.Context) {
	id := ctx.Param("id")
	category := controller.Service.FindById(id)
	ctx.JSON(200, gin.H{"data": category})
}

func (controller ControllerImpl) Save(ctx *gin.Context) {
	request := dto.AddCategoryDTO{}
	err := ctx.Bind(&request)
	utils.PanicIfError(err)
	controller.Service.Save(&request)
	ctx.JSON(200, gin.H{})
}

func (controller ControllerImpl) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var request dto.AddCategoryDTO
	ctx.Bind(&request)
	controller.Service.Update(id, &request)
	ctx.JSON(200, gin.H{})
}

func (controller ControllerImpl) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	controller.Service.Delete(id)
	ctx.JSON(200, gin.H{})
}

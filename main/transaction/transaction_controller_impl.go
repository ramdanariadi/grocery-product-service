package transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/ramdanariadi/grocery-product-service/main/exception"
	"github.com/ramdanariadi/grocery-product-service/main/transaction/dto"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"gorm.io/gorm"
)

type ControllerImpl struct {
	Service Service
}

func NewTransactionController(DB *gorm.DB) Controller {
	return ControllerImpl{Service: TransactionServiceImpl{DB: DB}}
}

func (controller ControllerImpl) Save(ctx *gin.Context) {
	var request dto.AddTransactionDTO
	ctx.ShouldBind(&request)
	userId, exists := ctx.Get("userId")
	if !exists {
		panic(exception.AuthenticationException{Message: "UNAUTHORIZED"})
	}
	controller.Service.save(&request, userId.(string))
	ctx.JSON(200, gin.H{})
}

func (controller ControllerImpl) Find(ctx *gin.Context) {
	var request dto.FindTransactionDTO
	err := ctx.ShouldBind(&request)
	utils.PanicIfError(err)
	transactionDTO := controller.Service.find(&request)
	ctx.JSON(200, gin.H{"data": transactionDTO})
}

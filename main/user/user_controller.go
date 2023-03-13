package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ramdanariadi/grocery-product-service/main/user/dto"
	"gorm.io/gorm"
)

type ControllerImpl struct {
	UserService Service
}

func NewUserController(db *gorm.DB) *ControllerImpl {
	return &ControllerImpl{UserService: NewUserService(db)}
}

func (controller *ControllerImpl) Register(ctx *gin.Context) {
	registerDTO := dto.RegisterDTO{}
	ctx.ShouldBind(&registerDTO)
	tokenDTO := controller.UserService.Register(&registerDTO)
	ctx.JSON(200, gin.H{"data": tokenDTO})
}

func (controller *ControllerImpl) Login(ctx *gin.Context) {
	loginDTO := dto.LoginDTO{}
	ctx.ShouldBind(&loginDTO)
	tokenDTO := controller.UserService.Login(&loginDTO)
	ctx.JSON(200, gin.H{"data": tokenDTO})
}

func (controller *ControllerImpl) Token(ctx *gin.Context) {
	tokenDTO := dto.TokenDTO{}
	ctx.ShouldBind(&tokenDTO)
	token := controller.UserService.Token(tokenDTO)
	ctx.JSON(200, gin.H{"data": token})
}

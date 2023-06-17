package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ramdanariadi/grocery-product-service/main/exception"
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

func (controller *ControllerImpl) Get(ctx *gin.Context) {
	value, exists := ctx.Get("userId")
	if !exists {
		panic(exception.AuthenticationException{Message: "FORBIDDEN"})
	}
	profileDTO := controller.UserService.Get(value.(string))
	ctx.JSON(200, gin.H{"data": profileDTO})
}

func (controller *ControllerImpl) Update(ctx *gin.Context) {
	updateProfileDTO := dto.ProfileDTO{}
	err := ctx.ShouldBind(&updateProfileDTO)
	if err != nil {
		panic(exception.ValidationException{Message: "BAD_REQUEST"})
	}

	value, exists := ctx.Get("userId")
	if !exists {
		panic(exception.AuthenticationException{Message: "FORBIDDEN"})
	}
	controller.UserService.Update(value.(string), &updateProfileDTO)
	ctx.JSON(200, gin.H{})
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

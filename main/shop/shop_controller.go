package shop

import "github.com/gin-gonic/gin"

type Controller interface {
	AddShop(ctx *gin.Context)
	EditShop(ctx *gin.Context)
	DeleteShop(ctx *gin.Context)
	GetShop(ctx *gin.Context)
}

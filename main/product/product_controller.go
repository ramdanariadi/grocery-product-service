package product

import "github.com/gin-gonic/gin"

type Controller interface {
	Save(ctx *gin.Context)
	FindById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	SetTopProduct(ctx *gin.Context)
	SetRecommendationProduct(ctx *gin.Context)
}

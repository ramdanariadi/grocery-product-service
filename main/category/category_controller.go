package category

import "github.com/gin-gonic/gin"

type CategoryController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

package wishlist

import "github.com/gin-gonic/gin"

type Controller interface {
	Store(ctx *gin.Context)
	Destroy(ctx *gin.Context)
	Find(ctx *gin.Context)
	FindByProductId(ctx *gin.Context)
}

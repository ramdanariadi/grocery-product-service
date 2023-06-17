package transaction

import "github.com/gin-gonic/gin"

type Controller interface {
	Save(ctx *gin.Context)
	Find(ctx *gin.Context)
}

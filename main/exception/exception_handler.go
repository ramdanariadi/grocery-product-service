package exception

import (
	"github.com/gin-gonic/gin"
)

func Handler(ctx *gin.Context, err any) {

	if errors, ok := err.(ValidationException); ok {
		ctx.AbortWithStatusJSON(400, gin.H{"message": errors.Message})
		return
	}

	if errors, ok := err.(AuthenticationException); ok {
		ctx.AbortWithStatusJSON(403, gin.H{"message": errors.Message})
		return
	}

	if err != nil {
		//if errors, ok := err.(error); ok {
		//	ctx.AbortWithStatusJSON(500, gin.H{"message": errors.Error()})
		//	return
		//}
		ctx.AbortWithStatusJSON(500, gin.H{"message": "INTERNAL_SERVER_ERROR"})
		return
	}
}

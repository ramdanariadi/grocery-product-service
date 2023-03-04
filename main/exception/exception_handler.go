package exception

import (
	"github.com/gin-gonic/gin"
)

func Handler(ctx *gin.Context, err any) {

	if errors, ok := err.(ValidationException); ok {
		ctx.JSON(400, gin.H{"message": errors.Message})
		return
	}

	if errors, ok := err.(AuthenticationException); ok {
		ctx.JSON(403, gin.H{"message": errors.Message})
		return
	}

	if err != nil {
		if errors, ok := err.(error); ok {
			ctx.JSON(500, gin.H{"message": errors.Error()})
			return
		}
		ctx.JSON(500, gin.H{"message": err})
		return
	}
}

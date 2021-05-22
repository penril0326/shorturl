package middleware

import (
	"github.com/gin-gonic/gin"
)

func RequestLimit(ctx *gin.Context) {
	limiter.Wait(ctx)

	ctx.Next()
}

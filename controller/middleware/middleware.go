package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func init() {
	postRate := rate.Every(10 * time.Millisecond)
	postLimitHandler = Limiter{
		rate.NewLimiter(postRate, POST_BUCKET_SIZE),
	}

	getRate := rate.Every(100 * time.Microsecond)
	getLimitHandler = Limiter{
		rate.NewLimiter(getRate, GET_BUCKET_SIZE),
	}
}

func PostRequestLimit(ctx *gin.Context) {
	postLimitHandler.Wait(ctx)

	ctx.Next()
}

func GetRequestLimit(ctx *gin.Context) {
	getLimitHandler.Wait(ctx)

	ctx.Next()
}

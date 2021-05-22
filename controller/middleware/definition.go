package middleware

import (
	"time"

	"golang.org/x/time/rate"
)

const (
	BUCKET_SIZE = 100
)

var r rate.Limit = rate.Every(time.Millisecond * 10)
var limiter *rate.Limiter = rate.NewLimiter(r, BUCKET_SIZE)

type Limiter struct {
	*rate.Limiter
}

var PostLimitHandler Limiter
var GetLimitHandler Limiter

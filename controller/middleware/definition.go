package middleware

import (
	"golang.org/x/time/rate"
)

const (
	POST_BUCKET_SIZE = 100
	GET_BUCKET_SIZE  = 100
)

type Limiter struct {
	*rate.Limiter
}

var postLimitHandler Limiter
var getLimitHandler Limiter

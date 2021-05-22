package main

import (
	"github.com/gin-gonic/gin"
	"github.com/penril0326/shorturl/controller/middleware"
	"github.com/penril0326/shorturl/controller/webapi"
	"github.com/penril0326/shorturl/cronjob"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	cronjob.Start()
}

func main() {
	defer cronjob.Stop()

	r := gin.Default()

	r.POST("/api/v1/urls", middleware.RequestLimit, webapi.CreateShort) // qps: 100/s
	r.DELETE("/api/v1/urls/:url_id", webapi.DeleteUrl)
	r.GET("/:url_id", webapi.Redirect) // qps: 10000/s

	r.Run()
}

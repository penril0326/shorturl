package router

import (
	"shorturl/controller/middleware"
	"shorturl/controller/webapi"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/api/v1/urls", middleware.PostRequestLimit, webapi.CreateShort)
	r.DELETE("/api/v1/urls/:url_id", webapi.DeleteUrl)
	r.GET("/:url_id", middleware.GetRequestLimit, webapi.Redirect)

	return r
}

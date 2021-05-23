package webapi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/penril0326/shorturl/model"
	"github.com/penril0326/shorturl/utils"
)

func CreateShort(ctx *gin.Context) {
	var input createShortUrlReq

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "invalid parameters")
		return
	}

	if input.Url == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "input url is empty")
		return
	}

	if utils.IsUrlValid(input.Url) == false {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "input url is invalid")
		return
	}

	short, err := model.Upsert(input.Url, input.ExpireAt)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "db error")
		return
	}

	ctx.JSON(http.StatusOK, createShortUrlResp{
		UrlID:    short,
		ShortUrl: fmt.Sprintf("http://%s/%s", "localhost:8080", short),
	})
}

func DeleteUrl(ctx *gin.Context) {
	urlID := ctx.Param("url_id")

	if urlID == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "url_id is empty")
		return
	}

	if err := model.DeleteByUrlID(urlID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "delete url failed")
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func Redirect(ctx *gin.Context) {
	urlID := ctx.Param("url_id")

	if urlID == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "url id is empty")
		return
	}

	urlInfo, err := model.GetUrlInfoByUrlID(urlID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "get origin url failed")
		return
	}

	if (urlInfo.OriginUrl == "") || (utils.IsT1AfterT2(time.Now(), urlInfo.ExpireAt) == true) {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, urlInfo.OriginUrl)
}

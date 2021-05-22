package webapi

import "time"

type createShortUrlReq struct {
	Url      string    `json:"url" binding:"required"`
	ExpireAt time.Time `json:"expire_at" binding:"required"`
}

type createShortUrlResp struct {
	UrlID    string `json:"id"`
	ShortUrl string `json:"short_url"`
}

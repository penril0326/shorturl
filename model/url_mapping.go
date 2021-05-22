package model

import (
	"errors"
	"time"

	"github.com/penril0326/shorturl/cache"
	"github.com/penril0326/shorturl/utils"
	"gorm.io/gorm"
)

const (
	TABLE_NAME_URL_MAPPING string = "url_mapping"
)

type UrlMapping struct {
	ID        int       `json:"id"`
	UrlID     string    `json:"url_id"`
	OriginUrl string    `json:"origin_url"`
	ExpireAt  time.Time `json:"expire_at"`
}

func (p *UrlMapping) TableName() string {
	return TABLE_NAME_URL_MAPPING
}

func Upsert(originUrl string, expireAt time.Time) (string, error) {
	urlInfo, err := GetUrlInfoByOriginUrl(originUrl)
	if err != nil {
		return "", err
	}

	tx := GetDB().Table(TABLE_NAME_URL_MAPPING).Begin()

	short := ""
	if urlInfo.OriginUrl != "" {
		if err := updateExpireTime(tx, urlInfo, expireAt); err != nil {
			tx.Rollback()
			return "", err
		}

		short = urlInfo.UrlID
	} else {
		urlId, err := insertUrlInfo(tx, originUrl, expireAt)
		if err != nil {
			tx.Rollback()
			return "", err
		}

		short = urlId
	}

	return short, nil
}

func updateExpireTime(tx *gorm.DB, urlInfo UrlMapping, expireAt time.Time) error {
	if expireAt.After(urlInfo.ExpireAt) == false {
		return nil
	}

	if result := tx.Model(&urlInfo).Update("expire_at", expireAt); result.Error != nil {
		return result.Error
	}

	return nil
}

func insertUrlInfo(tx *gorm.DB, originUrl string, expireAt time.Time) (string, error) {
	currentSeq, err := GetNextSequence(TABLE_NAME_URL_MAPPING)
	if err != nil {
		return "", err
	}

	urlID := utils.Base62Encode(int64(currentSeq + 1))
	new := UrlMapping{
		UrlID:     urlID,
		OriginUrl: originUrl,
		ExpireAt:  expireAt,
	}

	if result := tx.Create(&new); result.Error != nil {
		return "", result.Error
	}

	return urlID, nil
}

// func CreateShortUrl(originUrl string, expireAt time.Time) (string, error) {
// 	tx := GetDB().Table(TABLE_NAME_URL_MAPPING).Begin()

// 	currentSeq, err := GetNextSequence(TABLE_NAME_URL_MAPPING)
// 	if err != nil {
// 		tx.Rollback()
// 		return "", err
// 	}

// 	shortUrl := utils.Base62Encode(int64(currentSeq + 1))
// 	data := UrlMapping{
// 		UrlID:     shortUrl,
// 		OriginUrl: originUrl,
// 		ExpireAt:  expireAt,
// 	}

// 	if result := tx.Create(&data); result.Error != nil {
// 		tx.Rollback()
// 		return "", result.Error
// 	}

// 	tx.Commit()

// 	return shortUrl, nil
// }

func DeleteByUrlID(urlID string) error {
	tx := GetDB().Table(TABLE_NAME_URL_MAPPING).Begin()

	if tx == nil {
		return errors.New("get db transaction failed")
	}

	if result := tx.Where("url_id = ?", urlID).Delete(UrlMapping{}); result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	tx.Commit()

	if err := cache.DeleteAll(); err != nil {
		// log
	}

	return nil
}

func GetUrlInfoByOriginUrl(originUrl string) (UrlMapping, error) {
	db := GetDB().Table(TABLE_NAME_URL_MAPPING)

	var urlInfo UrlMapping
	if result := db.Select("origin_url, expire_at").Where("origin_url = ?", originUrl).Find(&urlInfo); result.Error != nil {
		return UrlMapping{}, result.Error
	}

	return urlInfo, nil
}

func GetUrlInfoByUrlID(urlID string) (UrlMapping, error) {
	db := GetDB().Table(TABLE_NAME_URL_MAPPING)

	var urlInfo UrlMapping
	if result := db.Select("*").Where("url_id = ?", urlID).Find(&urlInfo); result.Error != nil {
		return UrlMapping{}, result.Error
	}

	return urlInfo, nil
}

func DeleteExpired() error {
	tx := GetDB().Table(TABLE_NAME_URL_MAPPING).Begin()

	var expiredUrlID []string
	now := time.Now().UTC()
	if result := tx.Select("url_id").Where("expire_at <= ?", now).Find(&expiredUrlID); result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	if result := tx.Where("url_id IN (?)", expiredUrlID).Delete(UrlMapping{}); result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	tx.Commit()

	if err := cache.DeleteAll(); err != nil {
		// log
	}

	return nil
}
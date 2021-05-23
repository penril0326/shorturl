package model

import (
	"errors"
	"log"
	"shorturl/cache"
	"shorturl/utils"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
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

func Upsert(originUrl string, newExpireAt time.Time) (string, error) {
	urlInfo, err := GetUrlInfoByOriginUrl(originUrl)
	if err != nil {
		return "", err
	}

	tx := GetDB().Table(TABLE_NAME_URL_MAPPING).Begin()

	short := ""
	if urlInfo.OriginUrl != "" {
		if err := updateExpireTime(tx, urlInfo, newExpireAt); err != nil {
			tx.Rollback()
			return "", err
		}

		short = urlInfo.UrlID
	} else {
		urlId, err := insertUrlInfo(tx, originUrl, newExpireAt)
		if err != nil {
			tx.Rollback()
			return "", err
		}

		short = urlId
	}

	tx.Commit()

	return short, nil
}

func updateExpireTime(tx *gorm.DB, urlInfo UrlMapping, newExpireAt time.Time) error {
	if utils.IsT1AfterT2(newExpireAt, time.Now()) == false {
		return errors.New("expire time is invalid")
	} else if (utils.IsT1AfterT2(newExpireAt, urlInfo.ExpireAt) == false) &&
		(utils.IsT1AfterT2(newExpireAt, time.Now()) == true) {
		return nil
	} else {
		if result := tx.Model(&urlInfo).Update("expire_at", newExpireAt); result.Error != nil {
			return result.Error
		}

		if err := cache.DeleteKey(urlInfo.UrlID); (err != nil) && (err != memcache.ErrCacheMiss) {
			log.Println("Delete cache key error. Key: ", err.Error())
		}

		return nil
	}
}

func insertUrlInfo(tx *gorm.DB, originUrl string, newExpireAt time.Time) (string, error) {
	if utils.IsT1AfterT2(time.Now(), newExpireAt) == true {
		return "", errors.New("expire time is invalid")
	}

	currentSeq, err := GetCurrentSequence(TABLE_NAME_URL_MAPPING)
	if err != nil {
		return "", err
	}

	urlID := utils.Base62Encode(int64(currentSeq + 1))
	new := UrlMapping{
		UrlID:     urlID,
		OriginUrl: originUrl,
		ExpireAt:  newExpireAt,
	}

	if result := tx.Create(&new); result.Error != nil {
		return "", result.Error
	}

	if err := cache.Increase(cache.KEY_SEQUENCE); err != nil {
		log.Println("Increase cache value failed. Key: ", cache.KEY_SEQUENCE, ", Error: ", err.Error())
	}

	return urlID, nil
}

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

	if err := cache.DeleteKey(urlID); err != nil {
		log.Println("Delete key failed. Key: ", urlID, ", Error: ", err.Error())
	}

	return nil
}

func GetUrlInfoByOriginUrl(originUrl string) (UrlMapping, error) {
	db := GetDB().Table(TABLE_NAME_URL_MAPPING)

	var urlInfo UrlMapping
	if result := db.Select("*").Where("origin_url = ?", originUrl).Find(&urlInfo); result.Error != nil {
		return UrlMapping{}, result.Error
	}

	return urlInfo, nil
}

func GetUrlInfoByUrlID(urlID string) (UrlMapping, error) {
	var urlInfo UrlMapping
	if err := cache.GetData(urlID, &urlInfo); err == nil {
		return urlInfo, nil
	}

	db := GetDB().Table(TABLE_NAME_URL_MAPPING)
	if result := db.Select("*").Where("url_id = ?", urlID).Find(&urlInfo); result.Error != nil {
		return UrlMapping{}, result.Error
	}

	if err := cache.Set(urlID, urlInfo, 1800); err != nil {
		log.Println("Set cache error. Key: ", urlID, ", Error: ", err.Error())
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
		log.Println("Delete all cache key failed. Error: ", err.Error())
	}

	return nil
}

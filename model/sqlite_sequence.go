package model

import (
	"log"

	"github.com/penril0326/shorturl/cache"
)

const (
	TABLE_NAME_SQLITE_SEQUENCE = "sqlite_sequence"
)

func GetCurrentSequence(tableName string) (int64, error) {
	var currentSeq int64
	if err := cache.GetData(cache.KEY_SEQUENCE, &currentSeq); err == nil {
		return currentSeq, nil
	}

	db := GetDB().Table(TABLE_NAME_SQLITE_SEQUENCE)

	if result := db.Select("seq").Where("name = ?", tableName).Find(&currentSeq); result.Error != nil {
		return 0, result.Error
	}

	if err := cache.Set(cache.KEY_SEQUENCE, currentSeq, 0); err != nil {
		log.Println("Set cache key failed. Key: ", cache.KEY_SEQUENCE, ", Error: ", err.Error())
	}

	return currentSeq, nil
}

package cache

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/bradfitz/gomemcache/memcache"
)

func init() {
	server := fmt.Sprintf("%s:%s", HOST, PORT)
	mc = memcache.New(server)
	if mc == nil {
		log.Panic("failed to connect memcache")
	}
}

func Set(key string, value interface{}, expiredSecond int32) error {
	byteData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = mc.Set(&memcache.Item{
		Key:        key,
		Value:      byteData,
		Expiration: expiredSecond,
	})

	if err != nil {
		return err
	}

	return nil
}

func GetData(key string, data interface{}) error {
	it, err := mc.Get(key)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(it.Value, data); err != nil {
		return err
	}

	return nil
}

func DeleteKey(key string) error {
	return mc.Delete(key)
}

func DeleteAll() error {
	return mc.DeleteAll()
}

func Increase(key string) error {
	if _, err := mc.Increment(key, 1); err != nil {
		return err
	}

	return nil
}

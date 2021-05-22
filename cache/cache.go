package cache

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/bradfitz/gomemcache/memcache"
)

func init() {
	server := fmt.Sprintf("%s:%s", HOST, PORT)
	mc := memcache.New(server)
	if mc == nil {
		log.Panic("failed to connect memcache")
	}

	var count uint64
	b, _ := json.Marshal(&count)
	mc.Add(&memcache.Item{
		Key:   KEY_SEQUENCE,
		Value: b,
	})
}

func Set(key string, value interface{}, expiredSecond int32) {
	byteData, _ := json.Marshal(value)

	mc.Set(&memcache.Item{
		Key:        key,
		Value:      byteData,
		Expiration: expiredSecond,
	})
}

func Get(key string) string {
	it, err := mc.Get(key)
	if err != nil {
		return ""
	}

	var value string
	json.Unmarshal(it.Value, &value)

	return value
}

func DeleteAll() error {
	return mc.DeleteAll()
}

func GetCurrentSequence() (uint64, error) {
	it, err := mc.Get(KEY_SEQUENCE)
	if err != nil {
		return 0, err
	}

	var sequence uint64
	if err := json.Unmarshal(it.Value, &sequence); err != nil {
		return 0, err
	}

	return sequence, nil
}

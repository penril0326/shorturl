package cache

import "github.com/bradfitz/gomemcache/memcache"

const (
	HOST string = "localhost"
	PORT string = "11211"

	KEY_SEQUENCE = "sequence"
)

var mc *memcache.Client

package utils

import (
	"net/url"
	"time"
)

var base62Code = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

func Base62Encode(num int64) string {
	result := ""
	if num < 0 {
		return result
	}

	length := int64(len(base62Code))
	for num > 0 {
		remainder := num % length
		result = base62Code[remainder] + result
		num = (num) / 62
	}

	return result
}

func IsT1BeforeT2(t1, t2 time.Time) bool {
	return t1.UTC().Before(t2.UTC())
}

func IsUrlValid(originUrl string) bool {
	u, err := url.Parse(originUrl)
	if err != nil {
		return false
	}

	if ((u.Scheme != "http") && (u.Scheme != "https")) ||
		(u.Host == "localhost:8080") {
		return false
	}

	return true
}

package utils

import (
	"Infracloud_Assignment/singletons"
	"hash"
	"strings"
)

var hasher hash.Hash

func Shorten(longUrl string) (shortUrl string) {
	hashCode := GenerateHash(longUrl)
	split1Slice := strings.SplitAfter(longUrl, "://")
	httpString := split1Slice[0]
	split2Slice := strings.SplitAfter(split1Slice[1], "/")
	domainString := split2Slice[0]
	shortUrl = httpString + domainString + string(hashCode)
	return shortUrl
}

func GenerateHash(longUrl string) []byte {
	hasher = singletons.GetMd5Instance()
	hashCode := hasher.Sum([]byte(longUrl))
	return hashCode
}

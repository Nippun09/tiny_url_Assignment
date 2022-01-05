package utils

import (
	"Infracloud_Assignment/singletons"
	"fmt"
	"hash"
	"strings"
)

var hasher hash.Hash32

func Shorten(longUrl string) (shortUrl string) {
	fmt.Printf("long url string :%s\n", longUrl)
	hashCode := GenerateHash(longUrl)
	split1Slice := strings.SplitAfter(longUrl, "://")
	httpString := split1Slice[0]
	split2Slice := strings.SplitAfter(split1Slice[1], "/")
	domainString := split2Slice[0]
	shortUrl = httpString + domainString + fmt.Sprint(hashCode)
	return shortUrl
}

func GenerateHash(longUrl string) uint32 {
	hasher = singletons.GetFnvHaserInstance()
	hasher.Write([]byte(longUrl))
	return hasher.Sum32()
}

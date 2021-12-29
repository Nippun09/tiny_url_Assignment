package singletons

import (
	"crypto/md5"
	"hash"
	"sync"
)

var onceMd5 sync.Once

var md5Instance hash.Hash

func GetMd5Instance() hash.Hash {
	onceFunc := func() {
		md5Instance = md5.New()
	}
	onceMd5.Do(onceFunc)

	return md5Instance
}

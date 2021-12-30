package singletons

import (
	"hash"
	"hash/fnv"
	"sync"
)

var onceFnv sync.Once

var fnvhasher hash.Hash32

func GetFnvHaserInstance() hash.Hash32 {
	onceFunc := func() {
		fnvhasher = fnv.New32()
	}
	onceFnv.Do(onceFunc)

	return fnvhasher
}

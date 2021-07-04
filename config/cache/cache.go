package cache

import (
	"runtime/debug"
	"time"
	"fmt"

	"app/GoSample/logger"

	"github.com/coocood/freecache"
)

var cache *freecache.Cache = loadCache()

func loadCache() *freecache.Cache {
	cacheSize := 100 * 1024 * 1024
	cache := freecache.NewCache(cacheSize)
	fmt.Println("Cache created with size:", cacheSize)
	debug.SetGCPercent(20)
	return cache
}

// expireDuration = 0 means no expiration
func Set(key string, value string, expireDuration int) {
	keyBytes := []byte(key)
	valueBytes := []byte(value)
	cache.Set(keyBytes, valueBytes, expireDuration)
}

func Get(key string) string {
	keyBytes := []byte(key)
	value, err := cache.Get(keyBytes) //if cache has not key, returns value as empty string
	if err != nil {
		logger.ErrorLog("Cache couldn't find the key:", key, "- Error:", err.Error())
	}
	return string(value)
}

func Delete(key string) {
	cache.Del([]byte(key))
}

func Reset() {
	cache.Clear()
}

func GetAvaregeAccessTime() int64 {
	averageAccessTime := cache.AverageAccessTime() / int64(time.Second)
	return averageAccessTime
}

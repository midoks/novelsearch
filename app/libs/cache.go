package libs

import (
	"github.com/astaxie/beego/cache"
	"time"
)

func GetCacheConnId() (cache.Cache, error) {
	bm, err := cache.NewCache("file",
		`{"CachePath":"./.cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":0}`)
	return bm, err
}

func GetCache(key string) interface{} {
	linkId, err := GetCacheConnId()
	if err == nil {
		return linkId.Get(key)
	}
	r := make(map[interface{}]interface{})
	return r
}

func SetCache(key string, val interface{}, timeout time.Duration) error {
	linkId, err := GetCacheConnId()
	if err == nil {
		return linkId.Put(key, val, timeout)
	}
	return nil
}

package libs

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"time"
)

func GetCacheConnId() (cache.Cache, error) {
	file_cache := beego.AppConfig.String("file_cache")
	info := fmt.Sprintf(`{"CachePath":".cache","FileSuffix":"%s","DirectoryLevel":2,"EmbedExpiry":0}`, file_cache)
	bm, err := cache.NewCache("file", info)
	return bm, err
}

func GetCache(key string) (interface{}, error) {
	linkId, err := GetCacheConnId()
	if err == nil {
		return linkId.Get(key), nil
	}
	return nil, errors.New("缓存资源创建失败")
}

func SetCache(key string, val interface{}, timeout time.Duration) error {
	linkId, err := GetCacheConnId()
	if err == nil {
		return linkId.Put(key, val, timeout)
	}
	return nil
}

func CacheIsExist(key string) bool {
	linkId, err := GetCacheConnId()
	if err == nil {
		return linkId.IsExist(key)
	}
	return false
}

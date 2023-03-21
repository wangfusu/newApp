package utils

import (
	"NewApp/global"
	goCache "github.com/patrickmn/go-cache"
	"time"
)

type GoCache struct {
	cache *goCache.Cache
}

func (gc *GoCache) NewFunc() *goCache.Cache {
	// 默认5分钟过期，每10分钟删除过期的项目
	return goCache.New(10*time.Minute, 15*time.Minute)
}
func InitCache() {
	var gc *GoCache
	global.ICache = gc.NewFunc()
}

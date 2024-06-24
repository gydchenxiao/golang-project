package initilization

import (
	"github.com/patrickmn/go-cache"
	"time"
	"xkginweb/global"
)

// 初始化本地缓存
func InitCache() {
	c := cache.New(5*time.Minute, 10*time.Minute) // 设置时间
	global.Cache = c
}

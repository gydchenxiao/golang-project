package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/patrickmn/go-cache"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
	"xkginweb/commons/parse"
)

var (
	Cache      *cache.Cache
	Log        *zap.Logger
	SugarLog   *zap.SugaredLogger // -------------------------- 新增代码
	Lock       sync.RWMutex
	Yaml       map[string]interface{}
	Config     *parse.Config
	KSD_DB     *gorm.DB
	BlackCache local_cache.Cache
	REDIS      *redis.Client
)

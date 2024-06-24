package initilization

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"xkginweb/global"
)

func InitRedis() {
	redisCfg := global.Config.NoSQL.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Host,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.Dbname,   // use default DB // redis.Options 中的 DB 是一个 int 类型的变量
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("redis connect ping failed, err:", zap.Error(err))
	} else {
		fmt.Println("redis connect ping response:", zap.String("pong", pong))
		global.REDIS = client
	}
}

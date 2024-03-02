package initial

import (
	"go.uber.org/zap"
	"mxshop_api/user_api/global"
)
import "mxshop_api/user_api/data/redis_store"

func InitRedisStore() {
	global.RedisStore = redis_store.NewRedisStore(global.RedisCli)
	zap.S().Infof("InitRedisStore ok.")
}

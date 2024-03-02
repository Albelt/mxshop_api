package redis_store

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

const (
	expireTime       = time.Minute * 3
	captchaKeyFormat = "captcha_redis_store_$^*_%s"
)

// base64Captcha.Store的redis实现
type RedisStore struct {
	cli *redis.Client
}

func (r *RedisStore) Set(id string, value string) error {
	key := fmt.Sprintf(captchaKeyFormat, id)

	return r.cli.Set(context.Background(), key, value, expireTime).Err()
}

func (r *RedisStore) Get(id string, clear bool) string {
	key := fmt.Sprintf(captchaKeyFormat, id)

	res, err := r.cli.Get(context.Background(), key).Result()
	if err != nil {
		return ""
	}

	if clear {
		r.cli.Del(context.Background(), key)
	}

	return res
}

func (r *RedisStore) Verify(id, answer string, clear bool) bool {
	v := r.Get(id, clear)
	return v == answer
}

func NewRedisStore(rdsCli *redis.Client) *RedisStore {
	return &RedisStore{cli: rdsCli}
}

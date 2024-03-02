package global

import (
	"github.com/go-redis/redis/v8"
	consul_api "github.com/hashicorp/consul/api"
	"mxshop_api/oss_api/config"
)

var (
	Config    *config.Config
	ConsulCli *consul_api.Client
	RedisCli  *redis.Client
)

package global

import (
	pb "albelt.top/mxshop_protos/albelt/user_srv/go"
	"github.com/go-redis/redis/v8"
	consul_api "github.com/hashicorp/consul/api"
	"github.com/mojocn/base64Captcha"
	"mxshop_api/user_api/config"
	"mxshop_api/user_api/repo"
)

var (
	UserSrvCli pb.UserServiceClient
	RedisCli   *redis.Client
	RedisStore base64Captcha.Store
	Sms        repo.SmsRepo
	Config     *config.Config
	ConsulCli  *consul_api.Client
)

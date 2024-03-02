package global

import (
	pb "albelt.top/mxshop_protos/albelt/good_srv/go"
	"github.com/go-redis/redis/v8"
	consul_api "github.com/hashicorp/consul/api"
	"github.com/opentracing/opentracing-go"
	"mxshop_api/goods_api/config"
)

var (
	GoodsSrvCli pb.GoodsServiceClient
	RedisCli    *redis.Client
	Config      *config.Config
	ConsulCli   *consul_api.Client
	Tracer      opentracing.Tracer
)

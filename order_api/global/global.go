package global

import (
	pb "albelt.top/mxshop_protos/albelt/order_srv/go"
	"github.com/go-redis/redis/v8"
	consul_api "github.com/hashicorp/consul/api"
	"github.com/opentracing/opentracing-go"
	"mxshop_api/order_api/config"
)

var (
	RedisCli    *redis.Client
	Config      *config.Config
	ConsulCli   *consul_api.Client
	OrderSrvCli pb.OrderServiceClient
	Tracer      opentracing.Tracer
)

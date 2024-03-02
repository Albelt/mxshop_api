package initial

import (
	pb "albelt.top/mxshop_protos/albelt/good_srv/go"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mxshop_api/common/middleware/grpc_tracing"
	"mxshop_api/goods_api/global"
)

const (
	// %s:consul地址  %s:服务名称
	//grpcConsulUrlFormat = "consul://ubuntu-learn:8500/user-srv?wait=14s"
	grpcConsulUrlFormat = "consul://%s/%s?wait=14s"
)

func InitGoodsSrvCli(tracer opentracing.Tracer) {
	// 通过consul获取grpc服务的地址并配置客户端负载均衡策略
	grpcConsulUrl := fmt.Sprintf(grpcConsulUrlFormat, global.Config.Consul.Addr, global.Config.GoodsSrv.Name)

	conn, err := grpc.Dial(
		grpcConsulUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		// 使用tracing中间件
		grpc.WithUnaryInterceptor(
			grpc_tracing.OpenTracingClientInterceptor(tracer),
		),
	)
	if err != nil {
		panic(err)
	}

	global.GoodsSrvCli = pb.NewGoodsServiceClient(conn)
	zap.S().Info("InitGoodsSrvCli ok.")
}

package initial

import (
	pb "albelt.top/mxshop_protos/albelt/user_srv/go"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mxshop_api/user_api/global"
)

const (
	// %s:consul地址  %s:服务名称
	//grpcConsulUrlFormat = "consul://ubuntu-learn:8500/user-srv?wait=14s"
	grpcConsulUrlFormat = "consul://%s/%s?wait=14s"
)

func InitUserSrvClit() {
	// 通过consul获取grpc服务的地址并配置客户端负载均衡策略
	grpcConsulUrl := fmt.Sprintf(grpcConsulUrlFormat, global.Config.Consul.Addr, global.Config.UserSrv.Name)

	conn, err := grpc.Dial(
		grpcConsulUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		panic(err)
	}

	global.UserSrvCli = pb.NewUserServiceClient(conn)
	zap.S().Info("InitUserSrvClit ok.")
}

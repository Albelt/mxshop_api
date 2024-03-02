package initial

import (
	"fmt"
	consul_api "github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"mxshop_api/order_api/global"
)

var (
	serviceId string
)

// 初始化consul客户端
func InitConsul() {
	consulCfg := consul_api.DefaultConfig()
	consulCfg.Address = global.Config.Consul.Addr

	consulCli, err := consul_api.NewClient(consulCfg)
	if err != nil {
		panic(err)
	}

	global.ConsulCli = consulCli

	zap.S().Info("InitConsul ok.")
}

func RegisterConsul() {
	// 注册服务到consul
	externalIp := global.Config.Server.ExternalIp
	port := int(global.Config.Server.Port)
	serviceId = uuid.NewV4().String()
	srv := &consul_api.AgentServiceRegistration{
		ID:      serviceId,
		Name:    global.Config.Server.Name,
		Tags:    global.Config.Server.Tags,
		Port:    port,
		Address: externalIp,
		Check: &consul_api.AgentServiceCheck{
			Interval:                       global.Config.Consul.HealthCheck.Interval,
			Timeout:                        global.Config.Consul.HealthCheck.Timeout,
			HTTP:                           fmt.Sprintf("http://%s:%d/api/v1/health", externalIp, port),
			DeregisterCriticalServiceAfter: global.Config.Consul.HealthCheck.Deregister,
		},
	}

	if err := global.ConsulCli.Agent().ServiceRegister(srv); err != nil {
		panic(err)
	}

	zap.S().Info("RegisterConsul ok.")
}

func DeRegisterConsul() {
	if err := global.ConsulCli.Agent().ServiceDeregister(serviceId); err != nil {
		zap.S().Errorf("DeRegisterConsul failed:%s", err.Error())
	} else {
		zap.S().Info("DeRegisterConsul ok.")
	}
}

package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"mxshop_api/user_api/global"
	"mxshop_api/user_api/initial"
	"mxshop_api/user_api/middleware"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	ClOSE_TIMEOUT = time.Second * 0
)

func main() {
	// 初始化日志
	initial.InitLogger()

	// 初始化配置
	initial.InitConfig()

	// 初始化consul并注册服务
	initial.InitConsul()
	initial.RegisterConsul()

	// 初始化验证器
	initial.InitValidator()

	// 初始化中间件
	middleware.InitJwt()

	// 初始化UserSrvClit
	initial.InitUserSrvClit()

	// 初始化redis和RedisStore
	initial.InitRedis()
	initial.InitRedisStore()

	// 初始化sms
	initial.InitSms()

	// 初始化路由
	router := initial.InitRouters()

	// 启动服务器
	addr := fmt.Sprintf("%s:%d", global.Config.Server.Ip, global.Config.Server.Port)
	srv := http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		zap.S().Infof("Start server on:%s", addr)
		srv.ListenAndServe()
	}()

	// 主协程监听退出信息
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 从注册中心注销服务
	initial.DeRegisterConsul()

	// 优雅关闭
	ctx, cancel := context.WithTimeout(context.Background(), ClOSE_TIMEOUT)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.S().Errorf("Shutdown err:%s", err.Error())
	}

	select {
	case <-ctx.Done():
		zap.S().Infof("Shutdown timeout")
	}

	zap.S().Infof("Server Shutdown")
}

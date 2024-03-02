package initial

import (
	"github.com/gin-gonic/gin"
	"mxshop_api/common/middleware/gin_tracing"
	"mxshop_api/order_api/api"
	"mxshop_api/order_api/global"
	"os"
)

// 初始化路由
func InitRouters() *gin.Engine {
	r := gin.New()

	// 配置全局中间件
	r.Use(
		//日志
		gin.LoggerWithConfig(gin.LoggerConfig{
			Formatter: nil,
			Output:    os.Stdout,
			SkipPaths: []string{
				"/api/v1/health",
			},
		}),
		gin.Recovery(), //出错恢复
	)

	g := r.Group("/api/v1")

	// 健康检查
	g.GET("/health", api.HealthCheck)

	// 购物车相关
	cartG := g.Group("/cart_items")
	cartG.GET("/:user_id", api.GetCart)
	cartG.POST("/:user_id", api.AddItem)
	cartG.DELETE("/:user_id", api.DeleteItem)

	// 订单相关,使用tracing中间件
	orderG := g.Group("/order", gin_tracing.Tracing(global.Tracer))
	orderG.POST("/", api.CreateOrder)
	orderG.GET("/:order_id", api.GetOrder)

	return r
}

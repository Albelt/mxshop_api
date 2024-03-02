package initial

import (
	"github.com/gin-gonic/gin"
	"mxshop_api/common/middleware/gin_tracing"
	"mxshop_api/goods_api/api"
	"mxshop_api/goods_api/global"
	"mxshop_api/goods_api/middleware"
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

	// 商品相关,使用sentinel中间件,tracing中间件
	goodsG := g.Group("/goods", middleware.Sentinel(), gin_tracing.Tracing(global.Tracer))
	// 商品列表接口添加sentinel限流中间件
	goodsG.GET("/", api.ListGoods)
	goodsG.GET("/:goods_id", api.GetGoods)
	goodsG.POST("/", api.CreateGoods)
	goodsG.DELETE("/:goods_id", api.DeleteGoods)
	goodsG.PUT("/:goods_id", api.UpdateGoods)

	return r
}

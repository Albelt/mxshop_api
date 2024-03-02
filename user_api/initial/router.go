package initial

import (
	"github.com/gin-gonic/gin"
	"mxshop_api/user_api/api"
	"mxshop_api/user_api/middleware"
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
		gin.Recovery(),    //出错恢复
		middleware.Cors(), //跨域
	)

	g := r.Group("/api/v1")

	// 健康检查
	g.GET("/health", api.HealthCheck)

	// 用户相关
	userGroup := g.Group("/users")
	userGroup.GET("/", middleware.JwtAuth(), middleware.CheckAdmin(), api.ListUsers)
	userGroup.GET("/:user_id", middleware.JwtAuth(), api.GetUser)
	userGroup.POST("/login", api.LoginUser)
	userGroup.POST("/register", api.RegisterUser)
	userGroup.PUT("/:user_id", middleware.JwtAuth(), api.UpdateUser)

	// 其他
	utilGroup := g.Group("/utils")
	utilGroup.GET("/captcha", api.GetCaptcha)
	utilGroup.POST("/send_sms", api.SendSmsCode)

	return r
}

package initial

import (
	"github.com/gin-gonic/gin"
	"mxshop_api/oss_api/api"
	"mxshop_api/oss_api/middleware"
	"net/http"
	"os"
)

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

	// 配置静态文件
	r.LoadHTMLGlob("templates/*.html")
	// 配置静态文件夹路径 第一个参数是api，第二个是文件夹路径
	r.StaticFS("/static", http.Dir("static"))
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "OSS web直传",
		})
	})

	ossG := g.Group("/oss")
	ossG.GET("/token", api.Token)
	ossG.POST("/callback", api.HandlerRequest)

	return r
}

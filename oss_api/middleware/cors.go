package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		// Option请求和其他请求都需要返回这个header
		origin := c.GetHeader("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
		}

		if method == http.MethodOptions {
			// 设置跨域的相关header
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Access-Token, X-CSRF-Token, Authorization, Token, x-token")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")

			c.AbortWithStatus(http.StatusNoContent)
		}
	}
}

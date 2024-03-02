package middleware

import (
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/gin-gonic/gin"
	"net/http"
)

// sentinel中间件
func Sentinel() gin.HandlerFunc {
	return func(c *gin.Context) {
		resourceName := c.Request.Method + ":" + c.FullPath()
		entry, err := sentinel.Entry(resourceName,
			sentinel.WithResourceType(base.ResTypeWeb),
			sentinel.WithTrafficType(base.Inbound))

		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusTooManyRequests,
				gin.H{
					"path": resourceName,
					"err":  "too many requests",
				})
			return
		}

		defer entry.Exit()
		c.Next()
	}
}

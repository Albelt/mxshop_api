package gin_tracing

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"net/http"
)

const (
	tracingParentSpan = "parent-span"
)

// gin的tracing中间件
func Tracing(tracer opentracing.Tracer) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL.Path)

		span := tracer.StartSpan(
			name,
			opentracing.Tag{"user_agent", c.Request.UserAgent()},
			opentracing.Tag{"remote_address", c.Request.RemoteAddr},
		)
		defer span.Finish()

		c.Set(tracingParentSpan, span)
		c.Next()

		statusCode := c.Writer.Status()
		if statusCode != http.StatusOK {
			span.SetTag(string(ext.HTTPStatusCode), c.Writer.Status())
			c.Abort()
			return
		}
	}
}

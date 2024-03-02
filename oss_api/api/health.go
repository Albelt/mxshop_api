package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 健康检查
func HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

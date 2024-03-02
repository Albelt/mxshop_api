package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func handleParamErr(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": err.Error(),
	})
}

func handleGrpcErr(ctx *gin.Context, err error) {
	// 将grpc code转换成http code
	if s, ok := status.FromError(err); ok {
		switch s.Code() {
		case codes.AlreadyExists:
			ctx.JSON(http.StatusConflict, gin.H{
				"message": fmt.Sprintf("数据已存在:%s", s.Message()),
			})
		case codes.NotFound:
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("数据不存在:%s", s.Message()),
			})
		case codes.InvalidArgument:
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("参数错误:%s", s.Message()),
			})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    s.Code().String(),
				"message": s.Message(),
			})

		}
	}
}

func handleInternalErr(ctx *gin.Context, err error, note string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": fmt.Sprintf("%s:%s", note, err.Error()),
	})
}

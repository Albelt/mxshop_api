package middleware

import (
	pb "albelt.top/mxshop_protos/albelt/user_srv/go"
	"github.com/gin-gonic/gin"
	"mxshop_api/user_api/model"
	"net/http"
)

func CheckAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tmp, ok := c.Get(JwtContextKey)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": JwtContextKey + " not exists",
			})
			c.Abort()
			return
		}

		claim := tmp.(*model.CustomClaims)
		if claim.UserRole != int32(pb.UserRole_USER_ROLE_ADMIN) {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "无操作权限",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

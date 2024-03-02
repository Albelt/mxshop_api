package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"mxshop_api/user_api/global"
	"mxshop_api/user_api/model"
	"net/http"
	"time"
)

const (
	jwtTokenHeader = "access-token"
	JwtContextKey  = "jwt_claims"
)

var GlobalJwt *Jwt

func InitJwt() {
	var err error

	GlobalJwt, err = NewJwt()
	if err != nil {
		panic(err)
	}

	zap.S().Info("InitGlobalJwt ok.")
}

type Jwt struct {
	Key    []byte        // 签名秘钥
	Expire time.Duration // 过期时间(s)
	Issure string        // 签发人
}

func NewJwt() (*Jwt, error) {
	if global.Config.Jwt.Key == "" || global.Config.Jwt.Expire == "" {
		return nil, errors.New("invalid jwt config")
	}

	expire, err := time.ParseDuration(global.Config.Jwt.Expire)
	if err != nil {
		return nil, err
	}

	return &Jwt{
		Key:    []byte(global.Config.Jwt.Key),
		Expire: expire,
		Issure: global.Config.Jwt.Issure,
	}, nil
}

func (j *Jwt) GenToken(user *model.User) (string, error) {
	claims := model.CustomClaims{
		UserId:   int64(user.Id),
		UserName: user.NickName,
		UserRole: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.Issure,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.Expire)),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString(j.Key)
}

func (j *Jwt) ParseToken(tokenStr string) (*model.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &model.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.Key, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*model.CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token  string
			claims *model.CustomClaims
			err    error
		)

		// 从header中获取token
		token = c.GetHeader(jwtTokenHeader)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "token不存在",
			})
			c.Abort()
			return
		}

		// 解析并校验token
		if claims, err = GlobalJwt.ParseToken(token); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": fmt.Sprintf("无效的token:%s", err.Error()),
			})
			c.Abort()
			return
		}

		// 将解析后的claims对象放入上下文管理器
		c.Set(JwtContextKey, claims)
		c.Next()
	}
}

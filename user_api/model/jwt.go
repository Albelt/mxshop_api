package model

import (
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	UserRole int32  `json:"user_role"`
	jwt.RegisteredClaims
}

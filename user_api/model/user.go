package model

import (
	pb "albelt.top/mxshop_protos/albelt/user_srv/go"
	"time"
)

type User struct {
	Id       int32      `json:"id"`
	Mobile   string     `json:"mobile"`
	NickName string     `json:"nick_name"`
	Birthday *time.Time `json:"birthday"`
	Gender   int32      `json:"gender"`
	Role     int32      `json:"role"`
}

type ListUsersRequest struct {
	Page  int `form:"page" binding:"required,number"`
	Count int `form:"count" binding:"required,number"`
}

type ListUsersResponse struct {
	Total int32   `json:"total"`
	Users []*User `json:"users"`
}

type LoginUserRequest struct {
	Mobile   string `json:"mobile" binding:"required,mobile"`
	Password string `json:"password" binding:"required,min=6"`
	// 图形验证码
	CaptchaId     string `json:"captcha_id" binding:"required"`
	CaptchaAnswer string `json:"captcha_answer" binding:"required"`
}

type LoginUserResponse struct {
	Id       int32  `json:"id"`
	NickName string `json:"nick_name"`
	Token    string `json:"token"`
}

type RegisterUserRequest struct {
	Mobile   string `json:"mobile" binding:"required,mobile"`
	Password string `json:"password" binding:"required,min=6"`
	NickName string `json:"nick_name" binding:"required,min=2,max=64"`
	Code     string `json:"code" binding:"required"`
}

func NewUserFromPb(u *pb.User) *User {
	if u == nil {
		return nil
	}

	user := &User{
		Id:       u.Id,
		Mobile:   u.Mobile,
		NickName: u.NickName,
		Gender:   int32(u.Gender),
		Role:     u.Role,
	}

	if u.Birthday != nil {
		birthday := u.Birthday.AsTime()
		user.Birthday = &birthday
	}

	return user
}

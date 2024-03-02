package model

import pb "albelt.top/mxshop_protos/albelt/user_srv/go"

type GetCaptchaResponse struct {
	Id      string `json:"id"`
	Bs64Pic string `json:"bs_64_pic"`
}

type SendSmsCodeRequest struct {
	Mobile string     `json:"mobile" binding:"required,mobile"`
	Type   pb.SmsType `json:"type" binding:"required,oneof=1 2"`
}

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"mxshop_api/user_api/global"
	"mxshop_api/user_api/model"
	"net/http"
)

var driver = base64Captcha.NewDriverDigit(80, 200, 6, 0.7, 80)

// 获取图形验证码
func GetCaptcha(ctx *gin.Context) {
	var (
		err     error
		id      string
		bs64Pic string
	)

	captcha := base64Captcha.NewCaptcha(driver, global.RedisStore)

	if id, bs64Pic, err = captcha.Generate(); err != nil {
		handleInternalErr(ctx, err, "生成验证码错误")
		return
	}

	ctx.JSON(http.StatusOK, model.GetCaptchaResponse{
		Id:      id,
		Bs64Pic: bs64Pic,
	})
}

// 发送短信验证码
func SendSmsCode(ctx *gin.Context) {
	var (
		err  error
		req  model.SendSmsCodeRequest
		code string
	)

	// 校验参数
	if err = ctx.ShouldBindJSON(&req); err != nil {
		handleParamErr(ctx, err)
	}

	// 获取验证码
	if code, err = global.Sms.NewCode(req.Mobile); err != nil {
		handleInternalErr(ctx, err, "获取验证码失败")
	}

	// 短信发送验证码
	if err = global.Sms.SendSmsVerification(req.Mobile, code); err != nil {
		handleInternalErr(ctx, err, "发送验证码失败")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

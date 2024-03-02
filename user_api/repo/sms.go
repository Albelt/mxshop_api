package repo

type SmsRepo interface {
	// 发送短信验证码
	SendSmsVerification(mobile, code string) error
	// 生成并存储验证码
	NewCode(mobile string) (string, error)
	// 校验验证码
	VerifyCode(mobile, code string) (bool, error)
}

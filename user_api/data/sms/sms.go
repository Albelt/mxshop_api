package sms

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"math/rand"
	"mxshop_api/user_api/config"
	"mxshop_api/user_api/model"
	"time"
)

type Sms struct {
	cli *dysmsapi.Client
	cfg *config.Sms
	rds *redis.Client
}

func NewSms(cli *dysmsapi.Client, smsCfg *config.Sms, rds *redis.Client) *Sms {
	return &Sms{cli: cli, cfg: smsCfg, rds: rds}
}

// 发送验证码
func (s *Sms) SendSmsVerification(mobile, code string) error {
	params := model.SendVerificationParams{
		Code: code,
	}

	bytes, err := json.Marshal(params)
	if err != nil {
		return err
	}

	req := dysmsapi.SendSmsRequest{}
	req.SetSignName(s.cfg.Signature)
	req.SetTemplateCode(s.cfg.TemplateCode.Verification)
	req.SetPhoneNumbers(mobile)
	req.SetTemplateParam(string(bytes))

	zap.S().Infof("send request:%s", req.String())

	resp, _ := s.cli.SendSms(&req)
	if resp == nil || *resp.Body.Code != "OK" {
		zap.S().Errorf("阿里云短信发送失败:%s", resp.Body.String())
		return errors.New("SendSmsVerification err")
	}

	return nil
}

const (
	smsCodeKeyFormat = "sms_code_&*^_%s"
	smsCodeLen       = 6
	letters          = "01234567890"
	lettersLen       = len(letters)
	expireTime       = time.Minute * 3
)

// 生成并存储code
func (s *Sms) NewCode(mobile string) (string, error) {
	var (
		err  error
		code string
		key  string
	)

	key = fmt.Sprintf(smsCodeKeyFormat, mobile)
	code = randCode(smsCodeLen)

	err = s.rds.Set(context.Background(), key, code, expireTime).Err()
	if err != nil {
		return "", errors.New("NewCode store code failed.")
	}

	return code, nil
}

func randCode(length int) string {
	rand.Seed(time.Now().Unix())

	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = letters[rand.Intn(lettersLen)]
	}

	return string(bytes)
}

// 校验code
func (s *Sms) VerifyCode(mobile, code string) (bool, error) {
	var (
		err      error
		key      string
		realCode string
	)

	// 查询真实的code
	key = fmt.Sprintf(smsCodeKeyFormat, mobile)
	realCode, err = s.rds.Get(context.Background(), key).Result()
	if err != nil {
		return false, err
	}

	return realCode == code, nil
}

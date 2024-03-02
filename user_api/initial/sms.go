package initial

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"go.uber.org/zap"
	"mxshop_api/user_api/data/sms"
	"mxshop_api/user_api/global"
)

func InitSms() {
	config := openapi.Config{
		AccessKeyId:     &global.Config.Sms.AccessKeyId,
		AccessKeySecret: &global.Config.Sms.AccessKeySecret,
		Endpoint:        &global.Config.Sms.Endpoint,
	}

	client, err := dysmsapi.NewClient(&config)
	if err != nil {
		panic(err)
	}

	global.Sms = sms.NewSms(client, global.Config.Sms, global.RedisCli)
	zap.S().Info("InitSms ok.")
}

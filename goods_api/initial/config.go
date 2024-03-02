package initial

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mxshop_api/goods_api/config"
	"mxshop_api/goods_api/global"
)

const configFile = "config/config-dev.yaml"

// 使用viper读取配置文件
func InitConfig() {
	v := viper.New()
	v.SetConfigFile(configFile)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	var cfg config.Config
	if err := v.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	global.Config = &cfg

	zap.S().Infof("read config from %s, config:\n%s", configFile, global.Config.String())
}

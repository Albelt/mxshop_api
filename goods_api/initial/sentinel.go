package initial

import (
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"go.uber.org/zap"
)

// 初始化sentinel配置和限流规则
// 限流规则应该从文件/数据库中读取,而不是写死
func InitSentinelRateLimit() {
	if err := sentinel.InitDefault(); err != nil {
		panic(err)
	}

	_, err := flow.LoadRules([]*flow.Rule{
		{
			Resource:               "GET:/api/v1/goods/",
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
			Threshold:              1,
			StatIntervalInMs:       5000,
		},
	})
	if err != nil {
		panic(err)
	}

	zap.S().Info("InitSentinelRateLimit ok.")
}

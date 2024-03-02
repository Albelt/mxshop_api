package initial

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
	"io"
	"mxshop_api/order_api/global"
)

// 使用jaeger作为tracer
func InitTracer() (opentracing.Tracer, io.Closer) {
	serviceName := global.Config.Server.Name
	jaegerAgentAddr := global.Config.Tracing.JaegerAddr

	cfg := jaegercfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: jaegerAgentAddr,
		},
	}

	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(err)
	}
	zap.S().Info("InitTracer ok.")

	return tracer, closer
}

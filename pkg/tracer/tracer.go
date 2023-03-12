package tracer

import (
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

func NewJaegerTracer(serviceName, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: serviceName,
		//固定采样、对所有数据都进行采样
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		//是否启用loggingReporter、刷新缓冲区的频率、上报的Agent地址
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  agentHostPort,
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	//设置全局的Tracer对象
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}

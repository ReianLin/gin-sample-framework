package trace

import (
	"fmt"
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

type Config struct {
	ServiceName         string
	AgentHost           string
	AgentPort           string
	SamplerType         string
	SamplerParam        float64
	LogSpans            bool
	BufferFlushInterval time.Duration
}

func DefaultConfig() *Config {
	return &Config{
		ServiceName:         "gin-sample-framework",
		AgentHost:           "localhost",
		AgentPort:           "6831",
		SamplerType:         jaeger.SamplerTypeConst,
		SamplerParam:        1,
		LogSpans:            true,
		BufferFlushInterval: time.Second,
	}
}

func NewTracer(cfg *Config) (tracer opentracing.Tracer, closer io.Closer, err error) {
	if cfg == nil {
		cfg = DefaultConfig()
	}
	jaegerCfg := jaegercfg.Configuration{
		ServiceName: cfg.ServiceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  cfg.SamplerType,
			Param: cfg.SamplerParam,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            cfg.LogSpans,
			LocalAgentHostPort:  fmt.Sprintf("%s:%s", cfg.AgentHost, cfg.AgentPort),
			BufferFlushInterval: cfg.BufferFlushInterval,
		},
	}

	tracer, closer, err = jaegerCfg.NewTracer()
	if err != nil {
		return nil, nil, fmt.Errorf("cannot init Jaeger: %v", err)
	}

	opentracing.SetGlobalTracer(tracer)

	return tracer, closer, nil
}

package middleware

import (
	"gin-sample-framework/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

func Tracing(logger logger.Logger, tracer opentracing.Tracer) gin.HandlerFunc {
	return func(c *gin.Context) {
		spanCtx, _ := tracer.Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(c.Request.Header),
		)

		span := tracer.StartSpan(
			c.Request.URL.Path,
			opentracing.ChildOf(spanCtx),
		)
		defer span.Finish()

		// if jaegerSpan, ok := span.Context().(jaeger.SpanContext); ok {
		// 	logger.Info("request started", jaegerSpan.TraceID().String())
		// }

		c.Request = c.Request.WithContext(
			opentracing.ContextWithSpan(c.Request.Context(), span),
		)
		c.Next()
	}
}

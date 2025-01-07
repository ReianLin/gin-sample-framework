package trace

import (
	"context"

	"github.com/opentracing/opentracing-go"
)

func Trace(c context.Context, name string, fn func(context.Context) error) error {
	span, ctx := opentracing.StartSpanFromContext(c, name)
	defer span.Finish()
	return fn(ctx)
}

func TraceResult[T any](c context.Context, name string, fn func(context.Context) T) T {

	span, ctx := opentracing.StartSpanFromContext(c, name)
	defer span.Finish()

	return fn(ctx)
}

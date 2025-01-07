package model

import "context"

type Token struct {
	UserId   int
	UserRole string
}

var tokenDataKey struct{}

func GetTokenData[T Token](ctx context.Context) T {
	return ctx.Value(tokenDataKey).(T)
}

func SetTokenData[T Token](ctx context.Context, data T) context.Context {
	return context.WithValue(ctx, tokenDataKey, data)
}

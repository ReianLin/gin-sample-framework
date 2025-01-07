package service

import (
	"context"
	"gin-sample-framework/pkg/logger"
	"gin-sample-framework/pkg/trace"
)

type CatService struct {
	logger logger.Logger
}

func NewCatService(logger logger.Logger) *CatService {
	return &CatService{
		logger: logger,
	}
}

func (s *CatService) Action(ctx context.Context) string {
	return trace.TraceResult(ctx, "CatService.Action", func(ctx context.Context) string {
		s.logger.Info("cat is walking")
		return "cat walk"
	})
}

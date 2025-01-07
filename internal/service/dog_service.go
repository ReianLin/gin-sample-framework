package service

import (
	"context"
	"gin-sample-framework/pkg/logger"
	"gin-sample-framework/pkg/trace"
	"time"
)

type DogService struct {
	logger logger.Logger
}

func NewDogService(logger logger.Logger) *DogService {
	return &DogService{
		logger: logger,
	}
}

func (s *DogService) Action(ctx context.Context) string {
	return trace.TraceResult(ctx, "DogService.Action", func(ctx context.Context) string {
		s.logger.Info("dog is running")
		time.Sleep(5 * time.Second)
		return "dog run"
	})
}

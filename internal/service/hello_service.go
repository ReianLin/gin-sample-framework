package service

import (
	"context"
	"gin-sample-framework/internal/repository"
	"gin-sample-framework/pkg/logger"
	"gin-sample-framework/pkg/trace"
)

type HelloService struct {
	logger logger.Logger
	repo   *repository.HelloRepository
}

func NewHelloService(logger logger.Logger, repo *repository.HelloRepository) *HelloService {
	return &HelloService{
		logger: logger,
		repo:   repo,
	}
}

func (s *HelloService) SayHello(ctx context.Context) string {
	return trace.TraceResult(ctx, "Service.SayHello", func(ctx context.Context) string {
		s.logger.Info("service: saying hello")
		hello, _ := s.repo.GetHello()
		return hello
	})
}

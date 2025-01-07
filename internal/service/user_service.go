package service

import (
	"context"
	"gin-sample-framework/internal/model"
	"gin-sample-framework/internal/model/entity"
	"gin-sample-framework/internal/repository"
	"gin-sample-framework/pkg/logger"

	"go.uber.org/zap"
)

type UserService struct {
	logger   logger.Logger
	userRepo *repository.UserRepository
}

func NewUserService(logger logger.Logger, userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user model.UserCreateReq) (resp model.UserCreateResp, err error) {
	s.logger.Info("user create", zap.Any("user", user))
	userEntity := &entity.User{
		Username: user.Username,
		Password: user.Password,
	}
	err = s.userRepo.Create(ctx, userEntity)
	if err != nil {
		return resp, err
	}
	resp.UserId = userEntity.ID
	return resp, nil
}

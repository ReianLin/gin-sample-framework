package service

import (
	"context"
	"gin-sample-framework/internal/model"
	"gin-sample-framework/internal/model/entity"
	"gin-sample-framework/internal/repository"
	"gin-sample-framework/pkg/logger"
)

type UserService struct {
	logger      logger.Logger
	userRepo    *repository.UserRepository
	roleService *RoleService
}

func NewUserService(logger logger.Logger, userRepo *repository.UserRepository, roleService *RoleService) *UserService {
	return &UserService{
		userRepo:    userRepo,
		roleService: roleService,
		logger:      logger,
	}
}

func (s *UserService) Create(ctx context.Context, user model.UserCreateRequest) (resp model.UserResponse, err error) {
	userEntity := &entity.User{
		Username: user.Username,
		Password: user.Password,
	}
	err = s.userRepo.Create(ctx, userEntity, user.RoleIDs)
	if err != nil {
		return resp, err
	}
	resp.UserID = userEntity.UserID
	return resp, nil
}

func (s *UserService) GetAll(ctx context.Context) (users []*entity.User, err error) {
	users, err = s.userRepo.List(ctx)
	return
}

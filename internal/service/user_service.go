package service

import (
	"context"
	"fmt"
	"gin-sample-framework/internal/entity"
	"gin-sample-framework/internal/model"
	"gin-sample-framework/internal/repository"
	"gin-sample-framework/pkg/logger"
)

type UserService struct {
	logger   logger.Logger
	userRepo *repository.UserRepository
}

func NewUserService(logger logger.Logger, userRepo *repository.UserRepository) *UserService {
	return &UserService{
		logger:   logger,
		userRepo: userRepo,
	}
}

func (s *UserService) Create(ctx context.Context, req *model.UserCreateReq) (*model.UserCreateResp, error) {

	user := &entity.User{
		Account: req.Account,
		Email:   req.Email,
		Name:    req.Name,
		// Password: utils.EncryptPassword(req.Password),
		Password: req.Password,
	}

	err := s.userRepo.Transaction(ctx, func(txCtx context.Context) error {
		if err := s.userRepo.Create(txCtx, user); err != nil {
			return fmt.Errorf("create user failed: %w", err)
		}

		if len(req.RoleIDs) > 0 {
			if err := s.userRepo.AssignRolesCreate(txCtx, user.UserID, req.RoleIDs); err != nil {
				return fmt.Errorf("assign roles failed: %w", err)
			}
		}
		return nil
	})

	if err != nil {
		s.logger.Error("create user failed", "error", err)
		return nil, err
	}

	return &model.UserCreateResp{
		UserId: user.UserID,
	}, nil
}

func (s *UserService) GetDetail(ctx context.Context, userID int) (*model.UserDetailResp, error) {
	detail, err := s.userRepo.GetDetail(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &model.UserDetailResp{
		UserRoleDetailDTO: *detail,
	}, nil
}

func (s *UserService) Update(ctx context.Context, req model.UserUpdateReq) error {
	err := s.userRepo.Transaction(ctx, func(tx context.Context) error {

		if err := s.userRepo.Update(tx, req); err != nil {
			return err
		}

		if err := s.userRepo.AssignRolesDelete(tx, req.UserID); err != nil {
			return err
		}
		if len(req.RoleIDs) > 0 {
			if err := s.userRepo.AssignRolesCreate(tx, req.UserID, req.RoleIDs); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		s.logger.Error("update user failed", "error", err)
		return err
	}
	return nil
}

func (s *UserService) Delete(ctx context.Context, userID int) error {
	if err := s.userRepo.Delete(ctx, userID); err != nil {
		s.logger.Error("delete user failed", "error", err)
		return err
	}
	return nil
}

func (s *UserService) List(ctx context.Context, page, limit int) (*model.UserDetailListResponse, error) {
	users, total, err := s.userRepo.GetDetailList(ctx)
	if err != nil {
		s.logger.Error("get user list failed", "error", err)
		return nil, err
	}

	return &model.UserDetailListResponse{
		Total: total,
		Items: users,
	}, nil
}

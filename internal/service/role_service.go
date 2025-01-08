package service

import (
	"context"
	"fmt"
	"gin-sample-framework/internal/entity"
	"gin-sample-framework/internal/model"
	"gin-sample-framework/internal/repository"
	"gin-sample-framework/pkg/logger"
)

type RoleService struct {
	logger   logger.Logger
	roleRepo *repository.RoleRepository
}

func NewRoleService(logger logger.Logger, roleRepo *repository.RoleRepository) *RoleService {
	return &RoleService{
		logger:   logger,
		roleRepo: roleRepo,
	}
}

func (s *RoleService) Create(ctx context.Context, req *model.RoleCreateRequest) (*model.RoleCreateResponse, error) {
	roleEntity := &entity.Role{
		Name:        req.Name,
		Description: req.Description,
	}

	err := s.roleRepo.Transaction(ctx, func(txCtx context.Context) error {
		role, err := s.roleRepo.Create(txCtx, roleEntity)
		if err != nil {
			return fmt.Errorf("create role failed: %w", err)
		}
		roleEntity = role

		if len(req.PermissionIDs) > 0 {
			if err := s.roleRepo.AssignPermissionsCreate(txCtx, role.RoleID, req.PermissionIDs); err != nil {
				return fmt.Errorf("assign permissions failed: %w", err)
			}
		}
		return nil
	})

	if err != nil {
		s.logger.Error("create role failed", "error", err)
		return nil, err
	}

	return &model.RoleCreateResponse{
		RoleID: roleEntity.RoleID,
	}, nil
}

func (s *RoleService) Detail(ctx context.Context, roleID int) (*model.RoleDetailResponse, error) {
	detail, err := s.roleRepo.GetDetail(ctx, roleID)
	if err != nil {
		return nil, err
	}

	return &model.RoleDetailResponse{
		RolePermissionDetailDTO: *detail,
	}, nil
}

func (s *RoleService) Update(ctx context.Context, req *model.RoleUpdateRequest) error {
	err := s.roleRepo.Transaction(ctx, func(tx context.Context) error {
		if err := s.roleRepo.Update(tx, req); err != nil {
			return err
		}

		if err := s.roleRepo.AssignPermissionsDelete(tx, req.RoleID); err != nil {
			return err
		}

		if len(req.PermissionIDs) > 0 {
			if err := s.roleRepo.AssignPermissionsCreate(tx, req.RoleID, req.PermissionIDs); err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		s.logger.Error("update role failed", "error", err)
		return err
	}
	return nil
}

func (s *RoleService) Delete(ctx context.Context, roleID int) error {
	if err := s.roleRepo.Delete(ctx, roleID); err != nil {
		s.logger.Error("delete role failed", "error", err)
		return err
	}
	return nil
}

func (s *RoleService) List(ctx context.Context, page, limit int) (*model.RoleDetailListResponse, error) {
	roles, total, err := s.roleRepo.GetDetailList(ctx)
	if err != nil {
		s.logger.Error("get role list failed", "error", err)
		return nil, err
	}

	return &model.RoleDetailListResponse{
		Total: total,
		Items: roles,
	}, nil
}

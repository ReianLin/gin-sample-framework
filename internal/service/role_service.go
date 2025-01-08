package service

import (
	"context"
	"fmt"
	"gin-sample-framework/internal/model"
	"gin-sample-framework/internal/model/entity"
	"gin-sample-framework/internal/repository"
	"gin-sample-framework/pkg/logger"
)

type RoleService struct {
	logger         logger.Logger
	roleRepository *repository.RoleRepository
	userRepository *repository.UserRepository
}

func NewRoleService(logger logger.Logger, roleRepository *repository.RoleRepository, userRepository *repository.UserRepository) *RoleService {
	return &RoleService{
		logger:         logger,
		roleRepository: roleRepository,
		userRepository: userRepository,
	}
}

func (s *RoleService) Create(ctx context.Context, req *model.RoleCreateRequest) (resp model.RoleResponse, err error) {
	role, err := s.roleRepository.Create(ctx, &entity.Role{
		Name: req.Name,
	})
	if err != nil {
		return resp, err
	}
	resp.RoleID = role.RoleID
	return resp, nil
}

func (s *RoleService) GetAll(ctx context.Context) ([]*entity.Role, error) {
	roles, err := s.roleRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (s *RoleService) Get(ctx context.Context, id int) (*entity.Role, error) {
	role, err := s.roleRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if role == nil || role.RoleID == 0 {
		return nil, fmt.Errorf("role not found")
	}
	return role, nil
}

func (s *RoleService) Update(ctx context.Context, req *model.RoleUpdateRequest) (resp model.RoleResponse, err error) {
	if req.RoleID == 0 {
		return resp, fmt.Errorf("role id is required")
	}
	role, err := s.roleRepository.Get(ctx, req.RoleID)
	if err != nil {
		return resp, err
	}
	if role == nil || role.RoleID == 0 {
		return resp, fmt.Errorf("role not found")
	}
	if _, err := s.roleRepository.Update(ctx, req); err != nil {
		return resp, err
	}
	resp.RoleID = req.RoleID
	return resp, nil
}

func (s *RoleService) Delete(ctx context.Context, id int) error {
	if err := s.roleRepository.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

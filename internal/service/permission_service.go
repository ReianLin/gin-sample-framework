package service

import (
	"gin-sample-framework/internal/repository"
	"gin-sample-framework/pkg/logger"
)

type PermissionService struct {
	logger         logger.Logger
	permissionRepo *repository.PermissionRepository
}

func NewPermissionService(logger logger.Logger, permissionRepo *repository.PermissionRepository) *PermissionService {
	return &PermissionService{
		logger:         logger,
		permissionRepo: permissionRepo,
	}
}

package repository

import (
	"context"
	"gin-sample-framework/internal/db"
	"gin-sample-framework/internal/entity"
	"gin-sample-framework/pkg/logger"
)

type PermissionRepository struct {
	logger logger.Logger
	db     *db.DBProvider
}

func NewPermissionRepository(logger logger.Logger, db *db.DBProvider) *PermissionRepository {
	return &PermissionRepository{
		logger: logger,
		db:     db,
	}
}

func (r *PermissionRepository) GetAll(ctx context.Context) (permissions []*entity.Permission, err error) {
	err = r.db.WithDB(ctx).Model(&entity.Permission{}).Find(&permissions).Error
	return
}

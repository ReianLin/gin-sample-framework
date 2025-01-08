package repository

import (
	"context"
	"gin-sample-framework/internal/db"
	"gin-sample-framework/internal/model"
	"gin-sample-framework/internal/model/entity"
	"gin-sample-framework/pkg/logger"
)

type RoleRepository struct {
	db     *db.DBProvider
	logger logger.Logger
}

func NewRoleRepository(logger logger.Logger, db *db.DBProvider) *RoleRepository {
	return &RoleRepository{
		db:     db,
		logger: logger,
	}
}

func (r *RoleRepository) Create(ctx context.Context, req *entity.Role) (*entity.Role, error) {
	return nil, nil
}

func (r *RoleRepository) Get(ctx context.Context, id int) (role *entity.Role, err error) {
	err = r.db.WithDB(ctx).Model(&entity.Role{}).Where("role_id = ?", id).First(&role).Error
	return
}

func (r *RoleRepository) Update(ctx context.Context, req *model.RoleUpdateRequest) (*entity.Role, error) {
	var param = make(map[string]interface{})
	if req.Name != nil {
		param["name"] = req.Name
	}
	if err := r.db.WithDB(ctx).Model(&entity.Role{}).Where("role_id = ?", req.RoleID).Updates(param).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *RoleRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithDB(ctx).Where("role_id = ?", id).Delete(&entity.Role{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *RoleRepository) GetAll(ctx context.Context) (roles []*entity.Role, err error) {
	err = r.db.WithDB(ctx).Model(&entity.Role{}).Find(&roles).Error
	return
}

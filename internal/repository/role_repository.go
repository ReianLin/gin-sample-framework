package repository

import (
	"context"
	"gin-sample-framework/internal/db"
	"gin-sample-framework/internal/entity"
	"gin-sample-framework/internal/model"
	"gin-sample-framework/pkg/logger"

	"gorm.io/gorm"
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

func (r *RoleRepository) Transaction(ctx context.Context, fn func(txCtx context.Context) error) error {
	return r.db.WithDB(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(context.WithValue(ctx, db.GlobalDBProviderKey, tx))
	})
}

func (r *RoleRepository) Create(ctx context.Context, role *entity.Role) (*entity.Role, error) {
	err := r.db.WithDB(ctx).Model(&entity.Role{}).Create(&role).Error
	return role, err
}

func (r *RoleRepository) GetByID(ctx context.Context, id int) (*entity.Role, error) {
	var role entity.Role
	err := r.db.WithDB(ctx).Model(&entity.Role{}).Where("role_id = ?", id).First(&role).Error
	return &role, err
}

func (r *RoleRepository) Update(ctx context.Context, req *model.RoleUpdateRequest) error {
	params := make(map[string]interface{})
	if req.Name != nil {
		params["name"] = *req.Name
	}
	if req.Description != nil {
		params["description"] = *req.Description
	}
	return r.db.WithDB(ctx).Model(&entity.Role{}).Where("role_id = ?", req.RoleID).Updates(params).Error
}

func (r *RoleRepository) Delete(ctx context.Context, roleID int) error {
	return r.db.WithDB(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("role_id = ?", roleID).Delete(&entity.RolePermission{}).Error; err != nil {
			return err
		}
		if err := tx.Where("role_id = ?", roleID).Delete(&entity.Role{}).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *RoleRepository) GetAll(ctx context.Context) (roles []*entity.Role, err error) {
	err = r.db.WithDB(ctx).Model(&entity.Role{}).Find(&roles).Error
	return
}

func (r *RoleRepository) GetDetail(ctx context.Context, roleID int) (resp *model.RolePermissionDetailDTO, err error) {
	err = r.db.WithDB(ctx).Model(&entity.Role{}).Where("role_id = ?", roleID).Preload("Permissions").Find(&resp).Error
	return
}

func (r *RoleRepository) GetDetailList(ctx context.Context) (result []*model.RolePermissionDetailDTO, total int64, err error) {
	err = r.db.WithDB(ctx).Model(&entity.Role{}).Preload("Permissions").Find(&result).Count(&total).Error
	return
}

func (r *RoleRepository) AssignPermissionsCreate(ctx context.Context, roleID int, permissionIDs []int) error {
	rolePermissions := make([]entity.RolePermission, 0, len(permissionIDs))
	for _, permissionID := range permissionIDs {
		rolePermissions = append(rolePermissions, entity.RolePermission{
			RoleID:       roleID,
			PermissionID: permissionID,
		})
	}
	return r.db.WithDB(ctx).Create(&rolePermissions).Error
}

func (r *RoleRepository) AssignPermissionsDelete(ctx context.Context, roleID int) error {
	return r.db.WithDB(ctx).Where("role_id = ?", roleID).Delete(&entity.RolePermission{}).Error
}

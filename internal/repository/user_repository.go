package repository

import (
	"context"
	"gin-sample-framework/internal/db"
	"gin-sample-framework/internal/entity"
	"gin-sample-framework/internal/model"
	"gin-sample-framework/pkg/logger"

	"gorm.io/gorm"
)

type UserRepository struct {
	db     *db.DBProvider
	logger logger.Logger
}

func NewUserRepository(logger logger.Logger, db *db.DBProvider) *UserRepository {
	return &UserRepository{
		db:     db,
		logger: logger,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	return r.db.WithDB(ctx).Create(user).Error
}

func (r *UserRepository) GetDetail(ctx context.Context, userID int) (resp *model.UserRoleDetailDTO, err error) {
	err = r.db.WithDB(ctx).Model(&entity.User{}).Where("user_id = ?", userID).Preload("Roles").Find(&resp).Error
	return
}

func (r *UserRepository) GetDetailList(ctx context.Context) (result []*model.UserRoleDetailDTO, total int64, err error) {
	err = r.db.WithDB(ctx).Model(&entity.User{}).Find(&result).Count(&total).Error
	return
}

func (r *UserRepository) ChangePassword(ctx context.Context, userID int, password string) error {
	params := make(map[string]interface{})
	if password != "" {
		params["password"] = password
	}
	return r.db.WithDB(ctx).Model(&entity.User{}).Where("user_id = ?", userID).Updates(params).Error
}

func (r *UserRepository) Update(ctx context.Context, req model.UserUpdateReq) error {
	params := make(map[string]interface{})
	if req.Email != nil {
		params["email"] = *req.Email
	}

	if req.Name != nil {
		params["name"] = *req.Name
	}

	return r.db.WithDB(ctx).Model(&entity.User{}).Where("user_id = ?", req.UserID).Updates(params).Error
}

func (r *UserRepository) Delete(ctx context.Context, userID int) error {
	return r.db.WithDB(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ?", userID).Delete(&entity.UserRole{}).Error; err != nil {
			return err
		}
		if err := tx.Where("user_id = ?", userID).Delete(&entity.User{}).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *UserRepository) AssignRolesCreate(ctx context.Context, userID int, roleIDs []int) error {
	userRoles := make([]entity.UserRole, 0, len(roleIDs))
	for _, roleID := range roleIDs {
		userRoles = append(userRoles, entity.UserRole{
			UserID: userID,
			RoleID: roleID,
		})
	}
	return r.db.WithDB(ctx).Create(&userRoles).Error
}

func (r *UserRepository) AssignRolesDelete(ctx context.Context, userID int) error {
	return r.db.WithDB(ctx).Where("user_id = ?", userID).Delete(&entity.UserRole{}).Error
}

func (r *UserRepository) Transaction(ctx context.Context, fn func(txCtx context.Context) error) error {
	return r.db.WithDB(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(context.WithValue(ctx, db.GlobalDBProviderKey, tx))
	})
}

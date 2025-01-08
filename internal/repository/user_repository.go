package repository

import (
	"context"
	"fmt"
	"gin-sample-framework/internal/db"
	"gin-sample-framework/internal/model"
	"gin-sample-framework/internal/model/entity"
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

func (r *UserRepository) Create(ctx context.Context, user *entity.User, roleIDs []int) error {
	return r.db.WithDB(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		if len(roleIDs) > 0 {
			userRoles := make([]entity.UserRole, 0, len(roleIDs))
			for _, roleID := range roleIDs {
				userRoles = append(userRoles, entity.UserRole{
					UserID: user.UserID,
					RoleID: roleID,
				})
			}
			if err := tx.Create(&userRoles).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *UserRepository) Get(ctx context.Context, userID string) (*entity.User, []entity.Role, error) {
	var user entity.User
	var roles []entity.Role

	err := r.db.WithDB(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&user, "user_id = ?", userID).Error; err != nil {
			return err
		}
		var (
			userRolesTable = entity.UserRole{}.TableName()
			roleTable      = entity.Role{}.TableName()
		)
		if err := tx.Model(&roles).
			Joins(fmt.Sprintf("JOIN %s ON %s.role_id = %s.role_id", userRolesTable, userRolesTable, roleTable)).
			Where(fmt.Sprintf("%s.user_id = ?", userRolesTable), userID).
			Find(&roles).Error; err != nil {
			return err
		}
		return nil
	})

	return &user, roles, err
}

func (r *UserRepository) Update(ctx context.Context, req *model.UserUpdateRequest) error {
	return r.db.WithDB(ctx).Transaction(func(tx *gorm.DB) error {
		// 更新用户基本信息
		updates := make(map[string]interface{})
		if req.Username != nil {
			updates["username"] = *req.Username
		}
		if req.Password != nil {
			updates["password"] = *req.Password
		}

		if len(updates) > 0 {
			if err := tx.Model(&entity.User{}).Where("user_id = ?", req.UserID).Updates(updates).Error; err != nil {
				return err
			}
		}

		if req.RoleIDs != nil {
			if err := tx.Where("user_id = ?", req.UserID).Delete(&entity.UserRole{}).Error; err != nil {
				return err
			}

			if len(req.RoleIDs) > 0 {
				userRoles := make([]entity.UserRole, 0, len(req.RoleIDs))
				for _, roleID := range req.RoleIDs {
					userRoles = append(userRoles, entity.UserRole{
						UserID: req.UserID,
						RoleID: roleID,
					})
				}
				if err := tx.Create(&userRoles).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})
}

func (r *UserRepository) Delete(ctx context.Context, userID string) error {
	return r.db.WithDB(ctx).Transaction(func(tx *gorm.DB) error {
		// 删除用户角色关联
		if err := tx.Where("user_id = ?", userID).Delete(&entity.UserRole{}).Error; err != nil {
			return err
		}

		// 删除用户
		if err := tx.Where("user_id = ?", userID).Delete(&entity.User{}).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *UserRepository) List(ctx context.Context) ([]*entity.User, error) {
	var users []*entity.User
	err := r.db.WithDB(ctx).Find(&users).Error
	return users, err
}

func (r *UserRepository) GetUserRoles(ctx context.Context, userID string) ([]entity.Role, error) {
	var roles []entity.Role
	err := r.db.WithDB(ctx).Model(&entity.Role{}).
		Joins("JOIN user_roles ON user_roles.role_id = roles.role_id").
		Where("user_roles.user_id = ?", userID).
		Find(&roles).Error
	return roles, err
}

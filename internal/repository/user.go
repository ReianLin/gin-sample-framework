package repository

import (
	"context"
	"gin-sample-framework/internal/db"
	"gin-sample-framework/internal/model/entity"
	"gin-sample-framework/pkg/logger"
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

func (r *UserRepository) Create(ctx context.Context, user *entity.User) (err error) {
	err = r.db.WithDB(ctx).Model(&entity.User{}).Create(&user).Error
	return
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (result *entity.User, err error) {
	err = r.db.WithDB(ctx).Model(&entity.User{}).Where("id = ?", id).First(&result).Error
	return
}

func (r *UserRepository) Update(ctx context.Context, user *entity.User) (err error) {
	err = r.db.WithDB(ctx).Model(&entity.User{}).Updates(&user).Error
	return
}

func (r *UserRepository) Delete(ctx context.Context, id int) (err error) {
	err = r.db.WithDB(ctx).Model(&entity.User{}).Delete(&entity.User{}, id).Error
	return
}

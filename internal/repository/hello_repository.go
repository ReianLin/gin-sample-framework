package repository

import (
	"gin-sample-framework/internal/db"
	"gin-sample-framework/pkg/logger"
)

type HelloRepository struct {
	logger logger.Logger
	db     *db.DBProvider
}

func NewHelloRepository(logger logger.Logger, db *db.DBProvider) *HelloRepository {
	return &HelloRepository{
		logger: logger,
		db:     db,
	}
}

func (repo *HelloRepository) GetHello() (string, error) {
	return "Hello, World!", nil
}

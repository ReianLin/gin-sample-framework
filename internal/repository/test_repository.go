package repository

import (
	"context"
	"gin-sample-framework/internal/db"
	"gin-sample-framework/pkg/logger"
)

type TestRepository struct {
	logger logger.Logger
	db     *db.DBProvider
}

func NewTestRepository(logger logger.Logger, db *db.DBProvider) *TestRepository {
	return &TestRepository{
		logger: logger,
		db:     db,
	}
}

func (repo *TestRepository) GetTest(ctx context.Context) (string, error) {
	return "Test, World!", nil
}

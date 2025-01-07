package repository

import (
	"context"
	"gin-sample-framework/internal/db"
	"gin-sample-framework/internal/model/entity"
	"gin-sample-framework/pkg/logger"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, error) {
	// 创建 sqlmock
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	dialector := mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	return db, mock, nil
}

func TestUserRepository_Create(t *testing.T) {
	// 准备测试数据
	mockDB, mock, err := setupTestDB(t)
	assert.NoError(t, err)

	mockLogger := logger.NewZapLogger(zapcore.InfoLevel)
	repo := NewUserRepository(mockLogger, db.NewDBProvider(mockDB))

	user := &entity.User{
		Username: "test",
		Password: "password",
	}

	// 设置预期的 SQL 查询
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users`").
		WithArgs(user.Username, user.Password, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// 执行测试
	err = repo.Create(context.Background(), user)
	assert.NoError(t, err)

	// 验证所有期望的 SQL 都被执行
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetByID(t *testing.T) {
	mockDB, mock, err := setupTestDB(t)
	assert.NoError(t, err)

	mockLogger := logger.NewZapLogger(zapcore.InfoLevel)
	repo := NewUserRepository(mockLogger, db.NewDBProvider(mockDB))

	// 设置预期的查询结果
	rows := sqlmock.NewRows([]string{"id", "username", "password", "created_at", "updated_at"}).
		AddRow(1, "test", "test@example.com", time.Now(), time.Now())

	mock.ExpectQuery("^SELECT \\* FROM `users`").
		WithArgs(1).
		WillReturnRows(rows)

	// 执行测试
	user, err := repo.GetByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "test", user.Username)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_Update(t *testing.T) {
	mockDB, mock, err := setupTestDB(t)
	assert.NoError(t, err)

	mockLogger := logger.NewZapLogger(zapcore.InfoLevel)
	repo := NewUserRepository(mockLogger, db.NewDBProvider(mockDB))

	user := &entity.User{
		ID:       1,
		Username: "updated",
		Password: "newpassword",
	}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `users`").
		WithArgs(user.Username, user.Password, sqlmock.AnyArg(), 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Update(context.Background(), user)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_Delete(t *testing.T) {
	mockDB, mock, err := setupTestDB(t)
	assert.NoError(t, err)

	mockLogger := logger.NewZapLogger(zapcore.InfoLevel)
	repo := NewUserRepository(mockLogger, db.NewDBProvider(mockDB))

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM `users`").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Delete(context.Background(), 1)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

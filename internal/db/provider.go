package db

import (
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	globalProvider DBProvider
	globalKey      struct{}
)

type DBProvider struct {
	db *gorm.DB
}

func NewDBProvider(db *gorm.DB) *DBProvider {
	return &DBProvider{db: db}
}

func GetGlobalDBProvider() *DBProvider {
	return &globalProvider
}

func (d *DBProvider) WithDB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(globalKey).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db.WithContext(ctx)
}

func (d *DBProvider) DB() *gorm.DB {
	return d.db
}

type MysqlConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBName   string `json:"dbname"`
}

func (m *MysqlConfig) Dialector() gorm.Dialector {
	return mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true",
		m.Username, m.Password, m.Host, m.Port, m.DBName,
	))
}

func DBConnection(dial gorm.Dialector, opts ...func(*gorm.DB)) error {
	gormDB, err := gorm.Open(dial, &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}
	sqlDB, err := gormDB.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(500)
	if err = sqlDB.Ping(); err != nil {
		return err
	}

	globalProvider.db = gormDB
	return nil
}

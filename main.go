package main

import (
	"gin-sample-framework/config"
	_ "gin-sample-framework/docs"
	"gin-sample-framework/internal/db"
	"gin-sample-framework/internal/entity"
	"gin-sample-framework/internal/server"
	"gin-sample-framework/pkg/logger"
	"gin-sample-framework/pkg/trace"
	"os"
	"os/signal"
	"syscall"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

// @title           Gin Sample Framework API
// @version         1.0
// @description     This is a sample server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

func main() {

	// init signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// init config
	if err := config.Init(); err != nil {
		panic(err)
	}

	// init logger
	logs := logger.NewZapLogger(zapcore.InfoLevel)
	logger.SetGlobalLogger(logs)

	// init db
	var (
		dbConfig = db.MysqlConfig{
			Host:     config.Configuration.DB.MySQL.Host,
			Port:     config.Configuration.DB.MySQL.Port,
			Username: config.Configuration.DB.MySQL.Username,
			Password: config.Configuration.DB.MySQL.Password,
			DBName:   config.Configuration.DB.MySQL.DBName,
		}
	)
	if err := db.DBConnection(dbConfig.Dialector(), func(db *gorm.DB) {
		db.Use(dbresolver.Register(
			dbresolver.Config{
				Replicas: []gorm.Dialector{dbConfig.Dialector()},
			}))
	}); err != nil {
		logs.Error("failed to connect to database", zap.Error(err))
		panic(err)
	}

	db.GetGlobalDBProvider().DB().AutoMigrate(&entity.User{}, &entity.Role{}, &entity.UserRole{}, &entity.Permission{}, &entity.RolePermission{})

	// init tracer
	var tracer opentracing.Tracer
	if config.Configuration.Jaeger.Enabled {
		cfg := trace.DefaultConfig()
		cfg.ServiceName = config.Configuration.Name
		cfg.AgentHost = config.Configuration.Jaeger.Host
		cfg.AgentPort = config.Configuration.Jaeger.Port
		t, closer, err := trace.NewTracer(cfg)
		if err != nil {
			logs.Error("failed to create tracer", zap.Error(err))
			panic(err)
		}
		tracer = t
		defer closer.Close()
	}

	//init server
	if server := server.NewServer(logs, tracer); server != nil {
		go func() {
			if err := server.Run(); err != nil {
				logs.Error("failed to run server", zap.Error(err))
				panic(err)
			}
		}()
	}

	<-sigChan
}

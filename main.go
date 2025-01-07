package main

import (
	"gin-sample-framework/config"
	"gin-sample-framework/internal/db"
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

func main() {

	// init signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// init config
	if err := config.Init(); err != nil {
		panic(err)
	}

	// init logger
	logger := logger.NewZapLogger(zapcore.InfoLevel)

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
		logger.Error("failed to connect to database", zap.Error(err))
		panic(err)
	}

	// init tracer
	var tracer opentracing.Tracer
	if config.Configuration.Jaeger.Enabled {
		cfg := trace.DefaultConfig()
		cfg.ServiceName = config.Configuration.Name
		cfg.AgentHost = config.Configuration.Jaeger.Host
		cfg.AgentPort = config.Configuration.Jaeger.Port
		t, closer, err := trace.NewTracer(cfg)
		if err != nil {
			logger.Error("failed to create tracer", zap.Error(err))
			panic(err)
		}
		tracer = t
		defer closer.Close()
	}

	//init server
	if server := server.NewServer(logger, tracer); server != nil {
		go func() {
			if err := server.Run(); err != nil {
				logger.Error("failed to run server", zap.Error(err))
				panic(err)
			}
		}()
	}

	<-sigChan
}

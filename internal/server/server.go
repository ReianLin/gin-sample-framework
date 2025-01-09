package server

import (
	"fmt"
	"gin-sample-framework/config"
	"gin-sample-framework/pkg/logger"
	"gin-sample-framework/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

type Server struct {
	engine *gin.Engine
	logger logger.Logger
	tracer opentracing.Tracer
}

func NewServer(
	logger logger.Logger,
	tracer opentracing.Tracer,
) *Server {
	engine := gin.Default()
	return &Server{
		engine: engine,
		logger: logger,
		tracer: tracer,
	}
}

func (s *Server) Run() error {
	gin.SetMode(gin.DebugMode)
	s.InitMiddleware()
	initRoutes(s.engine)
	return s.engine.Run(fmt.Sprintf("%s:%d", "", config.Configuration.Server.Port))
}

func (s *Server) InitMiddleware() {
	s.engine.Use(middleware.Cors())
	// s.engine.Use(middleware.Auth())
	s.engine.Use(middleware.HasPermission())
	if s.tracer != nil {
		s.engine.Use(middleware.Tracing(s.logger, s.tracer))
	}
	s.engine.Use(middleware.Logger(s.logger))
}

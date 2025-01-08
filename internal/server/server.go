package server

import (
	"gin-sample-framework/internal/wire"
	"gin-sample-framework/pkg/logger"
	"gin-sample-framework/pkg/middleware"
	"os"

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
	s.InitRoutes()
	return s.engine.Run(":8080")
}

func (s *Server) InitMiddleware() {
	s.engine.Use(gin.RecoveryWithWriter(os.Stdout))
	s.engine.Use(middleware.Cors())
	// s.engine.Use(middleware.Auth())
	s.engine.Use(middleware.HasPermission())
	if s.tracer != nil {
		s.engine.Use(middleware.Tracing(s.logger, s.tracer))
	}
	s.engine.Use(middleware.Logger(s.logger))
}

type IRouter interface {
	Init(r *gin.RouterGroup)
}

func (s *Server) InitRoutes() {
	var routers = []IRouter{
		wire.BuildHelloController(),
		wire.BuildDogController(),
		wire.BuildCatController(),
	}
	for _, router := range routers {
		router.Init(s.engine.Group("/api/v1"))
	}
}

package server

import (
	"fmt"
	"gin-sample-framework/config"
	"gin-sample-framework/docs"
	"gin-sample-framework/internal/wire"
	"gin-sample-framework/pkg/logger"
	"gin-sample-framework/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	return s.engine.Run(fmt.Sprintf("%s:%d", "", config.Configuration.Server.Port))
}

func (s *Server) InitMiddleware() {
	// s.engine.Use(gin.RecoveryWithWriter(os.Stdout))
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
	docs.SwaggerInfo.Title = "W01-Swagger"
	docs.SwaggerInfo.Description = "W01-1.0"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", config.Configuration.Server.Port)
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	s.engine.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	var routers = []IRouter{
		// wire.BuildHelloController(),
		wire.BuildUserController(),
		wire.BuildRoleController(),
	}
	for _, router := range routers {
		router.Init(s.engine.Group("/api/v1"))
	}
}

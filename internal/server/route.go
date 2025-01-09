package server

import (
	"fmt"
	"gin-sample-framework/config"
	"gin-sample-framework/docs"
	"gin-sample-framework/internal/model"
	"gin-sample-framework/internal/wire"
	"gin-sample-framework/pkg/permission"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initRoutes(s *gin.Engine) {
	docs.SwaggerInfo.Title = "W01-Swagger"
	docs.SwaggerInfo.Description = "W01-1.0"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", config.Configuration.Server.Port)
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	s.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	group := s.Group("/api/v1")
	{

		//user
		{
			var groupName = "system/user"
			model.AppendMenuResourcesList("", "", groupName, "/"+groupName, permission.Create, permission.Delete, permission.Update, permission.Read)
			var (
				ctrl            = wire.BuildUserController()
				permissionGroup = permission.Permission.MakeGroup(groupName, groupName)
			)
			permissionGroup.Append(
				group.Group(groupName),
				permission.NewPerm("/create", http.MethodPost, permission.Create, ctrl.Create),
				permission.NewPerm("/detail/:user_id", http.MethodGet, permission.Read, ctrl.Detail),
				permission.NewPerm("/update", http.MethodPut, permission.Update, ctrl.Update),
				permission.NewPerm("/delete/:user_id", http.MethodDelete, permission.Delete, ctrl.Delete),
				permission.NewPerm("/list", http.MethodGet, permission.Read, ctrl.List),
			)
		}

		//role
		{
			var groupName = "system/role"
			model.AppendMenuResourcesList("", "", groupName, "/"+groupName, permission.Create, permission.Delete, permission.Update, permission.Read)
			var (
				ctrl            = wire.BuildRoleController()
				permissionGroup = permission.Permission.MakeGroup(groupName, groupName)
			)
			permissionGroup.Append(
				group.Group(groupName),
				permission.NewPerm("/create", http.MethodPost, permission.Create, ctrl.Create),
				permission.NewPerm("/detail/:role_id", http.MethodGet, permission.Read, ctrl.Detail),
				permission.NewPerm("/edit", http.MethodPut, permission.Update, ctrl.Edit),
				permission.NewPerm("/delete/:role_id", http.MethodDelete, permission.Delete, ctrl.Delete),
				permission.NewPerm("/list", http.MethodGet, permission.Read, ctrl.List),
			)

		}

	}
}

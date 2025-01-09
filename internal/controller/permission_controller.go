package controller

import (
	"gin-sample-framework/internal/service"
	"gin-sample-framework/pkg/logger"

	"github.com/gin-gonic/gin"
)

type PermissionController struct {
	baseController
	logger            logger.Logger
	permissionService *service.PermissionService
}

func NewPermissionController(logger logger.Logger, permissionService *service.PermissionService) *PermissionController {
	return &PermissionController{
		logger:            logger,
		permissionService: permissionService,
	}
}

// Create Permission
// @Tags         Permission
// @Summary      Create Permission
// @Description  Create a new permission
// @Accept       json
// @Produce      json
// @Param        permission  body  model.PermissionCreateRequest  true  "permission info"
// @Success      200  {object}  utils.GeneralResponseModel{data=model.PermissionCreateResponse}
// @Failure      400  {object}  utils.GeneralResponseModel
// @Router       /v1/system/permission/create [post]
func (ctrl *PermissionController) Create(c *gin.Context) {
	ctrl.Success(c, nil)

}

// List Permission
// @Tags         Permission
// @Summary      List Permission
// @Description  Get a list of permissions
// @Accept       json
// @Produce      json
// @Success      200  {object}  utils.GeneralResponseModel{data=[]model.Permission}
// @Failure      400  {object}  utils.GeneralResponseModel
// @Router       /v1/system/permission/list [get]
func (ctrl *PermissionController) List(c *gin.Context) {
	ctrl.Success(c, nil)

}

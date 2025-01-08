package controller

import (
	"gin-sample-framework/internal/service"
	"gin-sample-framework/pkg/logger"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	baseController
	logger      logger.Logger
	roleService *service.RoleService
	userService *service.UserService
}

func NewRoleController(logger logger.Logger, roleService *service.RoleService, userService *service.UserService) *RoleController {
	return &RoleController{
		baseController: baseController{
			Menu: menu{
				Name:  "role",
				Route: "/role",
			},
		},
		logger:      logger,
		roleService: roleService,
		userService: userService,
	}
}

// Create 创建角色
// @Tags         RoleController
// @Summary      Create
// @Description  Create
// @Accept       json
// @Produce      json
// @Param        request  body  model.RoleCreateRequest  true  "request"
// @Success      200  {object}  model.RoleCreateResponse
// @Failure      400  {object}  model.Error
// @Failure      404  {object}  model.Error
// @Router       /api/v1/role/create [post]
func (ctrl *RoleController) Create(c *gin.Context) {
}

// Get 获取角色
// @Tags         RoleController
// @Summary      Get
// @Description  Get
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Role
// @Failure      400  {object}  model.Error
// @Router       /api/v1/role/get/{id} [get]
func (ctrl *RoleController) Get(c *gin.Context) {

}

// Edit 编辑角色
// @Tags         RoleController
// @Summary      Edit
// @Description  Edit
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Role
// @Failure      400  {object}  model.Error
// @Router       /api/v1/role/edit/{id} [put]
func (ctrl *RoleController) Edit(c *gin.Context) {

}

// Delete 删除角色
// @Tags         RoleController
// @Summary      Delete
// @Description  Delete
// @Accept       json
// @Produce      json
// @Param        id  path  int  true  "id"
// @Success      200  {object}  model.Role
// @Failure      400  {object}  model.Error
// @Router       /api/v1/role/delete/{id} [delete]
func (ctrl *RoleController) Delete(c *gin.Context) {

}

// GetRoleList 获取角色列表
// @Tags         RoleController
// @Summary      GetRoleList
// @Description  GetRoleList
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.RoleListResponse
// @Router       /api/v1/role/list [get]
func (ctrl *RoleController) GetRoleList(c *gin.Context) {
}

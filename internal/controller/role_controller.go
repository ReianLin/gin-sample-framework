package controller

import (
	"gin-sample-framework/internal/model"
	"gin-sample-framework/internal/service"
	"gin-sample-framework/pkg/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	baseController
	logger      logger.Logger
	roleService *service.RoleService
}

func NewRoleController(logger logger.Logger, roleService *service.RoleService) *RoleController {
	return &RoleController{
		logger:      logger,
		roleService: roleService,
	}
}

// Create RoleController
// @Tags         Role
// @Summary      Create Role
// @Description  Create a new role
// @Accept       json
// @Produce      json
// @Param        role  body  model.RoleCreateRequest  true  "role info"
// @Success      200  {object}  utils.GeneralResponseModel{data=model.RoleCreateResponse}
// @Failure      400  {object}  utils.GeneralResponseModel
// @Router       /v1/system/role/create [post]
func (ctrl *RoleController) Create(c *gin.Context) {
	var req model.RoleCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.BadRequest(c, err)
		return
	}

	resp, err := ctrl.roleService.Create(c.Request.Context(), &req)
	if err != nil {
		ctrl.Error(c, err)
		return
	}

	ctrl.Success(c, resp)
}

// Get RoleDeta
// @Tags         Role
// @Summary      Get Role
// @Description  Get a role by ID
// @Accept       json
// @Produce      json
// @Param        role_id  path  int  true  "role_id"
// @Success      200  {object}  utils.GeneralResponseModel{data=model.RoleDetailResponse}  "response"
// @Failure      400  {object}  utils.GeneralResponseModel
// @Router       /v1/system/role/detail/{role_id} [get]
func (ctrl *RoleController) Detail(c *gin.Context) {
	roleID, err := strconv.Atoi(c.Param("role_id"))
	if err != nil {
		ctrl.BadRequest(c, err)
		return
	}

	resp, err := ctrl.roleService.Detail(c.Request.Context(), roleID)
	if err != nil {
		ctrl.Error(c, err)
		return
	}

	ctrl.Success(c, resp)
}

// Edit
// @Tags         Role
// @Summary      Edit
// @Description  Edit
// @Accept       json
// @Produce      json
// @Param        request  body  model.RoleUpdateRequest  true  "request"
// @Success      200  {object}  utils.GeneralResponseModel  "response"
// @Failure      400  {object}  utils.GeneralResponseModel
// @Router       /v1/system/role/edit [put]
func (ctrl *RoleController) Edit(c *gin.Context) {
	var req model.RoleUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.BadRequest(c, err)
		return
	}

	if err := ctrl.roleService.Update(c.Request.Context(), &req); err != nil {
		ctrl.Error(c, err)
		return
	}

	ctrl.Success(c, nil)
}

// Delete
// @Tags         Role
// @Summary      Delete
// @Description  Delete
// @Accept       json
// @Produce      json
// @Param        role_id  path  int  true  "role_id"
// @Success      200  {object}  utils.GeneralResponseModel  "response"
// @Failure      400  {object}  utils.GeneralResponseModel
// @Router       /v1/system/role/delete/{role_id} [delete]
func (ctrl *RoleController) Delete(c *gin.Context) {
	roleID, err := strconv.Atoi(c.Param("role_id"))
	if err != nil {
		ctrl.BadRequest(c, err)
		return
	}

	if err := ctrl.roleService.Delete(c.Request.Context(), roleID); err != nil {
		ctrl.Error(c, err)
		return
	}

	ctrl.Success(c, nil)
}

// List
// @Tags         Role
// @Summary      List
// @Description  List
// @Accept       json
// @Produce      json
// @Param        page   query  int  false  "page"
// @Param        limit  query  int  false  "limit"
// @Success      200  {object}  utils.GeneralResponseModel{data=model.RoleDetailListResponse}  "response"
// @Failure      400  {object}  utils.GeneralResponseModel
// @Router       /v1/system/role/list [get]
func (ctrl *RoleController) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	resp, err := ctrl.roleService.List(c.Request.Context(), page, limit)
	if err != nil {
		ctrl.Error(c, err)
		return
	}

	ctrl.Success(c, resp)
}

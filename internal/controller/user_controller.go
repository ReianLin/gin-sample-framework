package controller

import (
	"gin-sample-framework/internal/model"
	"gin-sample-framework/internal/service"
	"gin-sample-framework/pkg/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	baseController
	logger      logger.Logger
	userService *service.UserService
}

func NewUserController(logger logger.Logger, userService *service.UserService) *UserController {
	return &UserController{
		logger:      logger,
		userService: userService,
	}
}

// Create UserController
// @Tags         User
// @Summary      UserController
// @Description  UserController
// @Accept       json
// @Produce      json
// @Param        user  body  model.UserCreateReq  true  "user"
// @Success      200  {object}  utils.GeneralResponseModel{data=model.UserCreateResp}  "response"
// @Failure      400  {object}  utils.GeneralResponseModel
// @Router       /v1/system/user/create [post]
func (ctrl *UserController) Create(c *gin.Context) {
	var req model.UserCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.BadRequest(c, err)
		return
	}

	resp, err := ctrl.userService.Create(c.Request.Context(), &req)
	if err != nil {
		ctrl.Error(c, err)
		return
	}

	ctrl.Success(c, resp)
}

// Get UserController
// @Tags         User
// @Summary      Get
// @Description  Get
// @Accept       json
// @Produce      json
// @Param        user_id  path  string  true  "user_id"
// @Success      200  {object}  utils.GeneralResponseModel{data=model.UserDetailResp}  "response"
// @Failure      400  {object}  utils.GeneralResponseModel
// @Router       /v1/system/user/detail/{user_id} [get]
func (ctrl *UserController) Detail(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		ctrl.BadRequest(c, err)
		return
	}

	resp, err := ctrl.userService.GetDetail(c.Request.Context(), userID)
	if err != nil {
		ctrl.Error(c, err)
		return
	}

	ctrl.Success(c, resp)
}

// Update UserController
// @Tags         User
// @Summary      Update
// @Description  Update
// @Accept       json
// @Produce      json
// @Param        user  body  model.UserUpdateReq  true  "user"
// @Success      200  {object}  utils.GeneralResponseModel "response"
// @Failure      400  {object}  utils.GeneralResponseModel
// @Router       /v1/system/user/update [put]
func (ctrl *UserController) Update(c *gin.Context) {
	var req model.UserUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.BadRequest(c, err)
		return
	}

	if err := ctrl.userService.Update(c.Request.Context(), req); err != nil {
		ctrl.Error(c, err)
		return
	}

	ctrl.Success(c, nil)
}

// Delete UserController
// @Tags         User
// @Summary      Delete
// @Description  Delete
// @Accept       json
// @Produce      json
// @Param        user_id  path  string  true  "user_id"
// @Success      200  {object}  utils.GeneralResponseModel  "response"
// @Failure      400  {object}  utils.GeneralResponseModel
// @Router       /v1/system/user/delete/{user_id} [delete]
func (ctrl *UserController) Delete(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		ctrl.BadRequest(c, err)
		return
	}

	if err := ctrl.userService.Delete(c.Request.Context(), userID); err != nil {
		ctrl.Error(c, err)
		return
	}

	ctrl.Success(c, nil)
}

// List UserController
// @Tags         User
// @Summary      GetUserList
// @Description  GetUserList
// @Accept       json
// @Produce      json
// @Param        page  query  int  false  "page"
// @Param        limit  query  int  false  "limit"
// @Success      200  {object}  utils.GeneralResponseModel{data=model.UserDetailListResponse}  "response"
// @Failure      400  {object}  utils.GeneralResponseModel
// @Router       /v1/system/user/list [get]
func (ctrl *UserController) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	resp, err := ctrl.userService.List(c.Request.Context(), page, limit)
	if err != nil {
		ctrl.Error(c, err)
		return
	}

	ctrl.Success(c, resp)
}

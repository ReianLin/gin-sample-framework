package controller

import (
	"gin-sample-framework/internal/model"
	"gin-sample-framework/internal/service"
	"gin-sample-framework/pkg/logger"
	"gin-sample-framework/pkg/permission"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	baseController
	logger      logger.Logger
	userService *service.UserService
}

func NewUserController(logger logger.Logger, userService *service.UserService) *UserController {
	return &UserController{
		baseController: baseController{
			Menu: menu{
				Name:  "user",
				Route: "/user",
			},
		},
		logger:      logger,
		userService: userService,
	}
}

func (ctrl *UserController) Init(r *gin.RouterGroup) {
	model.AppendMenuResourcesList("", "", ctrl.Menu.Name, ctrl.Menu.Route, permission.Create, permission.Delete, permission.Update, permission.Read)
	group := r.Group(ctrl.Menu.Route)
	permission.Permission.MakeGroup(ctrl.Menu.Route, ctrl.Menu.Name).Append(group,
		permission.NewRoutePerm("/create", http.MethodPost, permission.Create, ctrl.Create),
	)
}

// Create UserController
// @Tags         User
// @Summary      UserController
// @Description  UserController
// @Accept       json
// @Produce      json
// @Param        user  body  model.UserCreateReq  true  "user"
// @Success      200  {object}  utils.Result{data=model.UserCreateResp}  "response"
// @Failure      400  {object}  utils.Result
// @Router       /api/v1/user/create [post]
func (ctrl *UserController) Create(c *gin.Context) {

}

// Get UserController
// @Tags         UserController
// @Summary      Get
// @Description  Get
// @Accept       json
// @Produce      json
// @Param        user_id  path  string  true  "user_id"
// @Success      200  {object}  model.User
// @Failure      400  {object}  model.Error
// @Router       /api/v1/user/get/{user_id} [get]
func (ctrl *UserController) Get(c *gin.Context) {

}

// Update UserController
// @Tags         UserController
// @Summary      Update
// @Description  Update
// @Accept       json
// @Produce      json
// @Param        user  body  model.UserUpdateReq  true  "user"
// @Success      200  {object}  model.User
// @Failure      400  {object}  model.Error
// @Router       /api/v1/user/update [put]
func (ctrl *UserController) Update(c *gin.Context) {

}

// Delete UserController
// @Tags         UserController
// @Summary      Delete
// @Description  Delete
// @Accept       json
// @Produce      json
// @Param        user_id  path  string  true  "user_id"
// @Success      200  {object}  model.User
// @Failure      400  {object}  model.Error
// @Router       /api/v1/user/delete/{user_id} [delete]
func (ctrl *UserController) Delete(c *gin.Context) {

}

// GetUserList UserController
// @Tags         UserController
// @Summary      GetUserList
// @Description  GetUserList
// @Accept       json
// @Produce      json
// @Param 	     page query int false "page for pagination" default(1)
// @Param        limit query int false "limit for pagination" default(10)
// @Success      200  {object}  model.UserListResponse
// @Failure      400  {object}  model.Error
// @Router       /api/v1/user/list [get]
func (ctrl *UserController) List(c *gin.Context) {

}

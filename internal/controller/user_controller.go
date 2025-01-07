package controller

import (
	"gin-sample-framework/internal/model"
	"gin-sample-framework/internal/service"
	"gin-sample-framework/pkg/logger"
	"gin-sample-framework/pkg/permission"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

// Walk UserController
// @Tags         UserController
// @Summary      UserController
// @Description  UserController
// @Accept       json
// @Produce      json
// @Param        user  body  model.UserCreateReq  true  "user"
// @Success      200  {object}  utils.Result{data=model.UserCreateResp}  "response"
// @Failure      400  {object}  utils.Result
// @Router       /api/v1/user/create [post]
func (ctrl *UserController) Create(c *gin.Context) {
	var (
		req model.UserCreateReq
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.logger.Error("user create error", zap.Error(err))
		return
	}
	if _, err := ctrl.userService.CreateUser(c.Request.Context(), req); err != nil {
		ctrl.logger.Error("user create error", zap.Error(err))
		return
	}
	ctrl.logger.Info("user create")
	c.JSON(200, gin.H{"message": "success"})
}

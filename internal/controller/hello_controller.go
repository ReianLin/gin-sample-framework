package controller

import (
	"gin-sample-framework/internal/model"
	"gin-sample-framework/internal/service"
	"gin-sample-framework/pkg/logger"
	"gin-sample-framework/pkg/permission"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloController struct {
	baseController
	logger  logger.Logger
	service *service.HelloService
}

func NewHelloController(
	logger logger.Logger,
	service *service.HelloService,
) *HelloController {
	return &HelloController{
		baseController: baseController{
			Menu: menu{
				Name:  "hello",
				Route: "/",
			},
		},
		logger:  logger,
		service: service,
	}
}

func (ctrl *HelloController) Init(r *gin.RouterGroup) {
	model.AppendMenuResourcesList("", "", ctrl.Menu.Name, ctrl.Menu.Route, permission.Create, permission.Delete, permission.Update, permission.Read)
	group := r.Group(ctrl.Menu.Route)
	permission.Permission.MakeGroup(ctrl.Menu.Route, ctrl.Menu.Name).Append(group,
		permission.NewRoutePerm("/hello", http.MethodGet, permission.Read, ctrl.Hello),
	)
}

// Hello TestHello
// @Tags         TestHello
// @Summary      TestHello
// @Description  TestHello
// @Accept       json
// @Produce      json
// @Success      200  {object}  utils.Result{data=string}  "response"
// @Failure      400  {object}  utils.Result
// @Router       /hello [post]
func (ctrl *HelloController) Hello(c *gin.Context) {

	result := ctrl.service.SayHello(c.Request.Context())

	c.JSON(200, gin.H{
		"message": result,
	})
}

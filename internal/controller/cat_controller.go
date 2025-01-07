package controller

import (
	"context"
	"gin-sample-framework/internal/model"
	"gin-sample-framework/internal/service"
	"gin-sample-framework/pkg/logger"
	"gin-sample-framework/pkg/permission"
	"gin-sample-framework/pkg/trace"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CatController struct {
	baseController
	logger  logger.Logger
	service service.IAnimalService
}

func NewCatController(logger logger.Logger, service service.IAnimalService) *CatController {
	return &CatController{
		baseController: baseController{
			Menu: menu{
				Name:  "cat",
				Route: "/cat",
			},
		},
		logger:  logger,
		service: service,
	}
}
func (ctrl *CatController) Init(r *gin.RouterGroup) {
	model.AppendMenuResourcesList("", "", ctrl.Menu.Name, ctrl.Menu.Route, permission.Create, permission.Delete, permission.Update, permission.Read)
	permission.Permission.MakeGroup(ctrl.Menu.Route, ctrl.Menu.Name).Append(r.Group(ctrl.Menu.Route),
		permission.NewRoutePerm("/walk", http.MethodGet, permission.Read, ctrl.Walk),
	)
}

// Walk TestCat
// @Tags         TestCat
// @Summary      TestCat
// @Description  TestCat
// @Accept       json
// @Produce      json
// @Success      200  {object}  utils.Result{data=string}  "response"
// @Failure      400  {object}  utils.Result
// @Router       /cat/walk [post]
func (ctrl *CatController) Walk(c *gin.Context) {
	result := trace.TraceResult(c.Request.Context(), "CatController.Walk", func(ctx context.Context) string {
		result := ctrl.service.Action(ctx)
		ctrl.logger.Info("cat walk")
		return result
	})
	c.JSON(200, gin.H{"message": result})
}

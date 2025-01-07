package controller

import (
	"context"
	"gin-sample-framework/internal/model"
	"gin-sample-framework/internal/service"
	"gin-sample-framework/pkg/logger"
	"gin-sample-framework/pkg/permission"
	"gin-sample-framework/pkg/trace"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type DogController struct {
	baseController
	logger  logger.Logger
	service service.IAnimalService
}

func NewDogController(logger logger.Logger, service service.IAnimalService) *DogController {
	return &DogController{
		baseController: baseController{
			Menu: menu{
				Name:  "dog",
				Route: "/dog",
			},
		},
		logger:  logger,
		service: service,
	}
}

func (ctrl *DogController) Init(r *gin.RouterGroup) {
	model.AppendMenuResourcesList("", "", ctrl.Menu.Name, ctrl.Menu.Route, permission.Create, permission.Delete, permission.Update, permission.Read)
	permission.Permission.MakeGroup(ctrl.Menu.Route, ctrl.Menu.Name).Append(r.Group(ctrl.Menu.Route),
		permission.NewRoutePerm("/run", http.MethodGet, permission.Read, ctrl.Run),
	)
}

// Run TestDog
// @Tags         TestDog
// @Summary      TestDog
// @Description  TestDog
// @Accept       json
// @Produce      json
// @Success      200  {object}  utils.Result{data=string}  "response"
// @Failure      400  {object}  utils.Result
// @Router       /dog/run [post]
func (c *DogController) Run(ctx *gin.Context) {
	result := trace.TraceResult(ctx.Request.Context(), "DogController.Run", func(ctx context.Context) string {
		result := c.service.Action(ctx)
		c.logger.Info("dog run")
		return result
	})

	time.Sleep(2 * time.Second)
	ctx.JSON(200, gin.H{"message": result})
}

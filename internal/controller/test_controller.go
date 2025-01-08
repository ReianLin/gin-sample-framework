package controller

import (
	"context"
	"gin-sample-framework/internal/model"
	"gin-sample-framework/internal/repository"
	"gin-sample-framework/internal/service"
	"gin-sample-framework/pkg/logger"
	"gin-sample-framework/pkg/permission"
	"gin-sample-framework/pkg/trace"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TestController struct {
	baseController
	logger     logger.Logger
	testRepo   *repository.TestRepository
	catService service.IAnimalService `wire:"catService"`
	dogService service.IAnimalService `wire:"dogService"`
}

func NewTestController(
	logger logger.Logger,
	testRepo *repository.TestRepository,
	catService service.IAnimalService,
	dogService service.IAnimalService,
) *TestController {
	return &TestController{
		baseController: baseController{
			Menu: menu{
				Name:  "Test",
				Route: "/test",
			},
		},
		logger:     logger,
		testRepo:   testRepo,
		catService: catService,
		dogService: dogService,
	}
}

func (ctrl *TestController) Init(r *gin.RouterGroup) {
	model.AppendMenuResourcesList("", "", ctrl.Menu.Name, ctrl.Menu.Route, permission.Create, permission.Delete, permission.Update, permission.Read)
	group := r.Group(ctrl.Menu.Route)
	permission.Permission.MakeGroup(ctrl.Menu.Route, ctrl.Menu.Name).Append(group,
		permission.NewRoutePerm("/hello", http.MethodGet, permission.Read, ctrl.Hello),
		permission.NewRoutePerm("/cat/walk", http.MethodGet, permission.Read, ctrl.Walk),
		permission.NewRoutePerm("/dog/run", http.MethodGet, permission.Read, ctrl.Run),
	)

}

// Hello TestHello
// @Tags         TestController
// @Summary      TestHello
// @Description  TestHello
// @Accept       json
// @Produce      json
// @Success      200  {object}  utils.GeneralResponseModel{data=string}  "response"
// @Failure      400  {object}  utils.GeneralResponseModel
// @Router       /v1/test/hello [get]
func (ctrl *TestController) Hello(c *gin.Context) {

	result, err := ctrl.testRepo.GetTest(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": result,
	})
}

// Walk TestCat
// @Tags         TestController
// @Summary      TestCat
// @Description  TestCat
// @Accept       json
// @Produce      json
// @Success      200  {object}  utils.GeneralResponseModel{data=string}  "response"
// @Failure      400  {object}  utils.GeneralResponseModel
// @Router       /v1/test/cat/walk [get]
func (ctrl *TestController) Walk(c *gin.Context) {
	result := trace.TraceResult(c.Request.Context(), "CatController.Walk", func(ctx context.Context) string {
		result := ctrl.catService.Action(ctx)
		ctrl.logger.Info("cat walk")
		return result
	})
	c.JSON(200, gin.H{"message": result})
}

// Run TestDog
// @Tags         TestController
// @Summary      TestDog
// @Description  TestDog
// @Accept       json
// @Produce      json
// @Success      200  {object}  utils.GeneralResponseModel{data=string}  "response"
// @Failure      400  {object}  utils.GeneralResponseModel
// @Router       /v1/test/dog/run [get]
func (ctrl *TestController) Run(ctx *gin.Context) {
	result := trace.TraceResult(ctx.Request.Context(), "DogController.Run", func(ctx context.Context) string {
		result := ctrl.dogService.Action(ctx)
		ctrl.logger.Info("dog run")
		return result
	})

	time.Sleep(2 * time.Second)
	ctx.JSON(200, gin.H{"message": result})
}

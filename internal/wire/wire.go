//go:build wireinject
// +build wireinject

package wire

import (
	"gin-sample-framework/internal/controller"
	"gin-sample-framework/internal/service"
	"gin-sample-framework/pkg/logger"

	"github.com/google/wire"
)

// Animal 相关的依赖注入，保留接口
func BuildDogController(logger logger.Logger) *controller.DogController {
	panic(wire.Build(
		service.NewDogService,
		wire.Bind(new(service.IAnimalService), new(*service.DogService)),
		controller.NewDogController,
	))
}

func BuildCatController(logger logger.Logger) *controller.CatController {
	panic(wire.Build(
		service.NewCatService,
		wire.Bind(new(service.IAnimalService), new(*service.CatService)),
		controller.NewCatController,
	))
}

func BuildHelloController(logger logger.Logger) *controller.HelloController {
	panic(wire.Build(
		HelloServiceSet,
		controller.NewHelloController,
	))
}

func BuildUserController(logger logger.Logger) *controller.UserController {
	panic(wire.Build(
		UserServiceSet,
		controller.NewUserController,
	))
}

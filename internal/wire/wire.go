//go:build wireinject
// +build wireinject

package wire

import (
	"gin-sample-framework/internal/controller"
	"gin-sample-framework/internal/service"
	"gin-sample-framework/pkg/logger"

	"github.com/google/wire"
)

func BuildDogController() *controller.DogController {
	panic(wire.Build(
		logger.GetGlobalLogger,
		service.NewDogService,
		wire.Bind(new(service.IAnimalService), new(*service.DogService)),
		controller.NewDogController,
	))
}

func BuildCatController() *controller.CatController {
	panic(wire.Build(
		logger.GetGlobalLogger,
		service.NewCatService,
		wire.Bind(new(service.IAnimalService), new(*service.CatService)),
		controller.NewCatController,
	))
}

func BuildHelloController() *controller.HelloController {
	panic(wire.Build(
		logger.GetGlobalLogger,
		DBProviderSet,
		HelloServiceSet,
		controller.NewHelloController,
	))
}

func BuildUserController() *controller.UserController {
	panic(wire.Build(
		logger.GetGlobalLogger,
		DBProviderSet,
		UserServiceSet,
		controller.NewUserController,
	))
}

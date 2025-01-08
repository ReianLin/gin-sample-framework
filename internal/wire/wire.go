//go:build wireinject
// +build wireinject

package wire

import (
	"gin-sample-framework/internal/controller"
	"gin-sample-framework/pkg/logger"

	"github.com/google/wire"
)

// func BuildTestController() *controller.TestController {
// 	panic(wire.Build(
// 		logger.GetGlobalLogger,
// 		DBProviderSet,
// 		repository.NewTestRepository,
// 		func() service.IAnimalService {
// 			panic(wire.Build(
// 				service.NewCatService,
// 				wire.Bind(new(service.IAnimalService), new(*service.CatService)),
// 			))
// 			return nil
// 		},
// 		func() service.IAnimalService {
// 			wire.Build(
// 				service.NewDogService,
// 				wire.Bind(new(service.IAnimalService), new(*service.DogService)),
// 			)
// 			return
// 		},
// 		controller.NewTestController,
// 	))
// }

func BuildUserController() *controller.UserController {
	panic(wire.Build(
		logger.GetGlobalLogger,
		DBProviderSet,
		UserServiceSet,
		controller.NewUserController,
	))
}

func BuildRoleController() *controller.RoleController {
	panic(wire.Build(
		logger.GetGlobalLogger,
		DBProviderSet,
		RoleServiceSet,
		controller.NewRoleController,
	))
}

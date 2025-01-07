package wire

import (
	"gin-sample-framework/internal/db"
	"gin-sample-framework/internal/repository"
	"gin-sample-framework/internal/service"

	"github.com/google/wire"
)

var (
	UserServiceSet = wire.NewSet(
		db.GetGlobalDBProvider,
		repository.NewUserRepository,
		service.NewUserService,
	)

	HelloServiceSet = wire.NewSet(
		db.GetGlobalDBProvider,
		repository.NewHelloRepository,
		service.NewHelloService,
	)
)

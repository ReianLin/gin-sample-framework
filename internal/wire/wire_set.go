package wire

import (
	"gin-sample-framework/internal/db"
	"gin-sample-framework/internal/repository"
	"gin-sample-framework/internal/service"

	"github.com/google/wire"
)

var DBProviderSet = wire.NewSet(
	db.GetGlobalDBProvider,
)

var (
	UserServiceSet = wire.NewSet(
		repository.NewUserRepository,
		RoleServiceSet,
		service.NewUserService,
	)

	HelloServiceSet = wire.NewSet(
		repository.NewHelloRepository,
		service.NewHelloService,
	)

	RoleServiceSet = wire.NewSet(
		repository.NewRoleRepository,
		service.NewRoleService,
	)
)

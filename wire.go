//+build wireinject

package main

import (
	"myproject/project/handler"
	"myproject/project/repository"
	"myproject/project/routes"
	"myproject/project/service"

	"github.com/google/wire"
)

func CreateUserConfig() (*routes.UserConfigRoutes, error) {
	wire.Build(
		repository.NewUserConfigRepository,
		service.NewUserConfigService,
		handler.NewUserConfigHandler,
		routes.NewUserConfigRoutes,
	)
	return &routes.UserConfigRoutes{}, nil
}

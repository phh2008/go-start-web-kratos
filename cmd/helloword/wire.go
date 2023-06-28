//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"helloword/internal/biz"
	"helloword/internal/conf"
	"helloword/internal/data"
	"helloword/internal/server"
	"helloword/internal/service"
	"helloword/pkg"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet,
		pkg.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		newApp,
	))
}

//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/scalent-io/shiprocket/pkg"
)

var ModuleSet = wire.NewSet(
	wire.Struct(new(DependencyOptions), "*"),
	NewShiprocketClient,

	pkg.NewShiprocketServiceImpl,
	wire.Bind(new(pkg.ShiprockertService), new(*pkg.ShiprocketServiceImpl)),
)

func initServer() (*ShiprocketClient, error) {
	panic(wire.Build(ModuleSet))
}

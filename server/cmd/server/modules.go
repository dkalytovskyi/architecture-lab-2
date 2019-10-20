//+build wireinject

package main

import (
	"github.com/dkalytovskyi/architecture-lab-2/server/plants"
	"github.com/google/wire"
)

func ComposeApiServer(port HttpPortNumber) (*GreenhouseApiServer, error) {
	wire.Build(
		NewDbConnection,
		plants.Householders,
		wire.Struct(new(GreenhouseApiServer), "Port", "PlantsHandler"),
	)
	return nil, nil
}

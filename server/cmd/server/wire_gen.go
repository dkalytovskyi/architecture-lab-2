// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/dkalytovskyi/architecture-lab-2/server/plants"
)

// Injectors from modules.go:

func ComposeApiServer(port HttpPortNumber) (*GreenhouseApiServer, error) {
	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}
	greenHouse := plants.NewGreenHouse(db)
	httpHandlerFunc := plants.HttpHandler(greenHouse)
	greenhouseApiServer := &GreenhouseApiServer{
		Port:          port,
		PlantsHandler: httpHandlerFunc,
	}
	return greenhouseApiServer, nil
}
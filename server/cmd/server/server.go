package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dkalytovskyi/architecture-lab-2/server/plants"
)

type HttpPortNumber int

type GreenhouseApiServer struct {
	Port HttpPortNumber

	PlantsHandler plants.HttpHandlerFunc

	server *http.Server
}

func (s *GreenhouseApiServer) Start() error {
	if s.PlantsHandler == nil {
		return fmt.Errorf("channels HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/plants", s.PlantsHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

func (s *GreenhouseApiServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}

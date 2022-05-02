package handlers

import (
	"fmt"
	"jellyfish/internal/libraries"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewHandler),
	fx.Provide(NewRoute),
)

type Route interface {
	Setup()
}

type route struct {
	echo    libraries.Echo
	handler Handler
}

func NewRoute(echo libraries.Echo, handler Handler) Route {
	return &route{
		echo:    echo,
		handler: handler,
	}
}

func (r *route) Setup() {
	fmt.Println("Setup")
	r.echo.Echo.GET("/health", r.handler.HealthCheck)
}

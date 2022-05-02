package handlers

import (
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
	netlify := r.echo.Echo.Group("/.netlify/functions/v1")
	{
		netlify.GET("/health", r.handler.HealthCheck)
	}
}

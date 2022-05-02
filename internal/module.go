package internal

import (
	"jellyfish/internal/handlers"
	"jellyfish/internal/libraries"

	"go.uber.org/fx"
)

func NewModule(bootstrap interface{}) fx.Option {
	var Module = fx.Options(
		handlers.Module,
		libraries.Module,
		fx.Invoke(bootstrap),
	)

	return Module
}

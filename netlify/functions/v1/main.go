package main

import (
	"context"
	"jellyfish/internal"
	"jellyfish/internal/handlers"
	"jellyfish/internal/libraries"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func bootstrap(
	lifecycle fx.Lifecycle,
	route handlers.Route,
	echo libraries.Echo,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					// echo.Echo.Validator = validator
					// echo.Echo.HTTPErrorHandler = libraries.ErrorHandler

					route.Setup()
					server := libraries.Route(echo.Echo)
					lambda.Start(server)
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				return nil
			},
		},
	)
}

func main() {
	godotenv.Load()

	var Module = internal.NewModule(bootstrap)
	fx.New(Module).Run()
}

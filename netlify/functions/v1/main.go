package main

import (
	"context"
	"jellyfish/internal/handlers"
	"jellyfish/internal/libraries"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

var Module = fx.Options(
	handlers.Module,
	libraries.Module,
	fx.Invoke(bootstrap),
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
	fx.New(Module).Run()
}

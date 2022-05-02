package main

import (
	"context"
	"fmt"
	"jellyfish/internal"
	"jellyfish/internal/handlers"
	"jellyfish/internal/libraries"
	"os"

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
					port, found := os.LookupEnv("PORT")

					if !found {
						port = "1323"
					}

					// echo.Echo.Validator = validator
					// echo.Echo.HTTPErrorHandler = libraries.ErrorHandler

					route.Setup()

					echo.Echo.Logger.Fatal(
						echo.Echo.Start(fmt.Sprintf(":%s", port)),
					)
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

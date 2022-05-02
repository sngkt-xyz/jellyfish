package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hey")
	})

	server := Route(e)

	lambda.Start(server)
}

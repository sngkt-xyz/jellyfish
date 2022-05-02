package libraries

import "github.com/labstack/echo/v4"

type Echo struct {
	Echo *echo.Echo
}

func NewEcho() Echo {
	e := echo.New()
	return Echo{Echo: e}
}

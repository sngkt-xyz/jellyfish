package handlers

import (
	"jellyfish/internal/constants"
	"jellyfish/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, models.NewResponse(http.StatusOK, constants.SuccessHealthCheck, models.HealthCheckResponse{
		Status: "OK",
	}))
}

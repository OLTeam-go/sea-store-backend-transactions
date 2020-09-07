package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *httpDelivery) Hello(c echo.Context) error {
	return c.JSON(http.StatusAccepted, "Hello from transaction service")
}

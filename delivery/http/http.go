package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//Response represent the response of the request
type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func (h *httpDelivery) Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello from transaction service")
}

func (h *httpDelivery) FetchBanks(c echo.Context) error {
	ctx := c.Request().Context()
	res, err := h.bankUsecase.Fetch(ctx)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status: http.StatusOK,
		Data:   res,
	})
}

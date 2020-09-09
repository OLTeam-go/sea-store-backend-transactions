package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
)

func getPageFromRequest(c echo.Context) (int, error) {
	qparam := c.QueryParam("page")
	page, err := strconv.Atoi(qparam)
	if err != nil {
		page = 1
	}
	if page <= 0 {
		page = 1
	}
	return page, nil
}

func (h *httpDelivery) Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello from transaction service")
}

func (h *httpDelivery) FetchBanks(c echo.Context) error {
	ctx := c.Request().Context()
	res, err := h.bankUsecase.Fetch(ctx)
	log.Println(res)
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

func (h *httpDelivery) GetActiveCartByCustomerID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	ctx := c.Request().Context()

	res, err := h.cartUsecase.GetActiveByCustomerID(ctx, id)

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

func (h *httpDelivery) FetchCartHistoryByCustomerID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	page, err := getPageFromRequest(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	ctx := c.Request().Context()

	res, err := h.cartUsecase.FetchHistoryByCustomerID(ctx, id, page)

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

func (h *httpDelivery) AddItemToCart(c echo.Context) error {
	var body CartItemRequest
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	c.Bind(&body)
	body.Quantity = 1
	ctx := c.Request().Context()

	err = h.cartItemUsecase.AddItemToCart(ctx, id, body.ItemID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status: http.StatusOK,
	})
}

func (h *httpDelivery) RemoveItemFromCart(c echo.Context) error {
	var body CartItemRequest
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	c.Bind(&body)
	ctx := c.Request().Context()

	err = h.cartItemUsecase.RemoveItemFromCart(ctx, id, body.ItemID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status: http.StatusOK,
	})
}

package http

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/OLTeam-go/sea-store-backend-transactions/enum"

	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
)

func getPageFromRequest(c echo.Context) (int, error) {
	qparam := c.QueryParam("page")
	page, err := strconv.Atoi(qparam)
	if err != nil {
		page = 0
	}
	if page < 0 {
		return 0, errors.New("page is invalid")
	}
	return page, nil
}

func getFilterFromRequest(c echo.Context) enum.TransactionFilterStatus {
	fparam := c.QueryParam("filter")
	switch fparam {
	case string(enum.TransactionFilterAccepted):
		return enum.TransactionFilterAccepted
	case string(enum.TransactionFilterAll):
		return enum.TransactionFilterAll
	case string(enum.TransactionFilterPending):
		return enum.TransactionFilterPending
	case string(enum.TransactionFilterRejected):
		return enum.TransactionFilterRejected
	default:
		return enum.TransactionFilterAll
	}
}

func (h *httpDelivery) Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello from transaction service")
}

// FetchBanks process request to fetch all available banks
// @Summary Endpoint to fetch all available banks
// @Accept json
// @Produce json
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /banks [get]
// @Success 200 {object} Response{data=[]models.Bank}
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

// GetActiveCartByCustomerID process request to get an active cart for a customer
// @Summary Endpoint to get active cart of an user
// @Accept json
// @Produce json
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /cart/customer/:id [get]
// @Success 200 {object} Response{data=models.Cart}
// @Param customer_id path string true "Customer ID"
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

// FetchCartHistoryByCustomerID process request to fetch cart history for a customer
// @Summary Endpoint to fetch cart history for a customer
// @Accept json
// @Produce json
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /cart/customer/history/{customer_id} [get]
// @Success 200 {object} Response{data=[]models.Cart}
// @Param customer_id path string true "Customer ID"
// @Param page query int false "page index default 1 (1 based index), omit means fetch all"
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

	return c.JSON(http.StatusOK, PaginationResponse{
		Status: http.StatusOK,
		Data:   res,
		Size:   len(res),
		Page:   page,
	})
}

// AddItemToCart process request to add an item to customer cart
// @Summary Endpoint to add an item to customer cart
// @Accept json
// @Produce json
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /cart/customer/add/{customer_id} [post]
// @Success 200 {object} Response
// @Param customer_id path string true "Customer ID"
// @Param default body CartItemRequest true "Cart Item REquest"
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

// AddItemToCart process request to add an item to customer cart
// @Summary Endpoint to add an item to customer cart
// @Accept json
// @Produce json
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /cart/customer/remove/{customer_id} [post]
// @Success 200 {object} Response
// @Param customer_id path string true "Customer ID"
// @Param default body CartItemRequest true "Cart Item Request"
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

// CheckoutCart process request to checkout a cart
// @Summary Endpoint to checkout a cart
// @Accept json
// @Produce json
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /transaction/checkout/{customer_id} [post]
// @Success 200 {object} Response
// @Param customer_id path string true "Customer ID"
// @Param default body CheckoutRequest true "Checkout Request"
func (h *httpDelivery) CheckoutCart(c echo.Context) error {
	customerID, err := uuid.Parse(c.Param("customer_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	var body CheckoutRequest
	err = c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	ctx := c.Request().Context()
	res, err := h.transactionUsecase.CreateTransaction(ctx, customerID, body.BankID, body.BankAccountNumber)
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

// AcceptTransaction process request to accept a transaction by admin
// @Summary Endpoint to accept a transaction by admin
// @Accept json
// @Produce json
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /transaction/accept/{transaction_id} [post]
// @Success 200 {object} Response
// @Param transaction_id path string true "Transaction ID"
func (h *httpDelivery) AcceptTransaction(c echo.Context) error {
	transactionID, err := uuid.Parse(c.Param("transaction_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	ctx := c.Request().Context()

	err = h.transactionUsecase.AcceptStatusTransaction(ctx, transactionID)
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

// RejectTransaction process request to reject a transaction by admin
// @Summary Endpoint to reject a transaction by admin
// @Accept json
// @Produce json
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /transaction/reject/{transaction_id} [post]
// @Success 200 {object} Response
// @Param transaction_id path string true "Transaction ID"
func (h *httpDelivery) RejectTransaction(c echo.Context) error {
	transactionID, err := uuid.Parse(c.Param("transaction_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	ctx := c.Request().Context()

	err = h.transactionUsecase.RejectStatusTransaction(ctx, transactionID)
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

// FetchTransactions process request to fetch transactions
// @Summary Endpoint to fetch transactions
// @Accept json
// @Produce json
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /transactions [get]
// @Success 200 {object} Response{data=[]models.Transaction}
// @Param page query int false "page index default 1 (1 based index), omit means fetch all"
// @Param filter query int false "Filter available = [all, rejected, pending, accepted] default all"
func (h *httpDelivery) FetchTransactions(c echo.Context) error {
	page, err := getPageFromRequest(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	filter := getFilterFromRequest(c)

	ctx := c.Request().Context()
	res, err := h.transactionUsecase.FetchTransactions(ctx, page, filter)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, PaginationResponse{
		Status: http.StatusOK,
		Data:   res,
		Page:   page,
		Size:   len(res),
	})
}

// FetchTransactionsByCustomerID process request to fetch transactions history for a customer
// @Summary Endpoint to fetch transactions
// @Accept json
// @Produce json
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /transactions/history/{customer_id} [get]
// @Success 200 {object} Response{data=[]models.Transaction}
// @Param customer_id path string true "customer uuid"
// @Param page query int false "page index default 1 (1 based index), omit means fetch all"
// @Param filter query int false "Filter available = [all, rejected, pending, accepted] default all"
func (h *httpDelivery) FetchTransactionsByCustomerID(c echo.Context) error {
	customerID, err := uuid.Parse(c.Param("customer_id"))
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
	filter := getFilterFromRequest(c)

	ctx := c.Request().Context()
	res, err := h.transactionUsecase.FetchTransactionsByCustomerID(ctx, page, customerID, filter)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, PaginationResponse{
		Status: http.StatusOK,
		Data:   res,
		Page:   page,
		Size:   len(res),
	})
}

// FetchRequestedItemsByMerchantID process request to fetch paid snapshot items for a merchant
// @Summary Endpoint to fetch paid snapshot items for a merchant
// @Accept json
// @Produce json
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /transactions/merchant/{merchant_id} [get]
// @Success 200 {object} Response{data=[]models.SnapshotCartItem}
// @Param merchant_id path string false "Merchant ID"
// @Param page query int false "page index default 1 (1 based index), omit means fetch all"
func (h *httpDelivery) FetchRequestedItemsByMerchantID(c echo.Context) error {
	merchantID, err := uuid.Parse(c.Param("merchant_id"))
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
	res, err := h.snapshotUsecase.FetchSnapshotCartItemsByMerchantID(ctx, page, merchantID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, PaginationResponse{
		Status: http.StatusOK,
		Data:   res,
		Page:   page,
		Size:   len(res),
	})
}

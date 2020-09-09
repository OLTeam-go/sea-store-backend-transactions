package http

import (
	dTransactions "github.com/OLTeam-go/sea-store-backend-transactions/delivery"
	"github.com/OLTeam-go/sea-store-backend-transactions/domain"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type httpDelivery struct {
	bankUsecase     domain.BankUsecase
	cartUsecase     domain.CartUsecase
	cartItemUsecase domain.CartItemUsecase
}

func cartGroup(e *echo.Echo, h *httpDelivery) {
	cart := e.Group("/api/cart")
	cart.GET("/customer/:id", h.GetActiveCartByCustomerID)
	cart.POST("/customer/:id/add", h.AddItemToCart)
	cart.POST("/customer/:id/remove", h.RemoveItemFromCart)
	cart.GET("/customer/history/:id", h.FetchCartHistoryByCustomerID)
}

func apiGroup(e *echo.Echo, h *httpDelivery) {
	api := e.Group("/api")
	api.GET("/banks", h.FetchBanks)
	api.GET("/docs/*", echoSwagger.WrapHandler)
	cartGroup(e, h)
}

// New function intialize delivery implementation
func New(e *echo.Echo, bu domain.BankUsecase, cu domain.CartUsecase, ciu domain.CartItemUsecase) dTransactions.Delivery {
	handler := &httpDelivery{
		bankUsecase:     bu,
		cartUsecase:     cu,
		cartItemUsecase: ciu,
	}
	e.GET("/", handler.Hello)
	apiGroup(e, handler)
	return handler
}

package http

import (
	dTransactions "github.com/OLTeam-go/sea-store-backend-transactions/delivery"
	"github.com/OLTeam-go/sea-store-backend-transactions/domain"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type httpDelivery struct {
	bankUsecase domain.BankUsecase
}

func apiGroup(e *echo.Echo, h *httpDelivery) {
	api := e.Group("/api")
	api.GET("/banks", h.FetchBanks)
	api.GET("/docs/*", echoSwagger.WrapHandler)
}

// New function intialize delivery implementation
func New(e *echo.Echo, bu domain.BankUsecase) dTransactions.Delivery {
	handler := &httpDelivery{
		bankUsecase: bu,
	}
	e.GET("/", handler.Hello)
	apiGroup(e, handler)
	return handler
}

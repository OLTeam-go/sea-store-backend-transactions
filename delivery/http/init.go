package http

import (
	dTransactions "github.com/OLTeam-go/sea-store-backend-transactions/delivery"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type httpDelivery struct {
}

// New function intialize delivery implementation
func New(e *echo.Echo) dTransactions.Delivery {
	handler := &httpDelivery{}
	e.GET("/", handler.Hello)
	api := e.Group("/api")
	api.GET("/docs/*", echoSwagger.WrapHandler)
	return handler
}

package delivery

import "github.com/labstack/echo/v4"

// Delivery contains method used for handling user request
type Delivery interface {
	Hello(c echo.Context) error
	FetchBanks(c echo.Context) error
	GetActiveCartByCustomerID(c echo.Context) error
	FetchCartHistoryByCustomerID(c echo.Context) error
	AddItemToCart(c echo.Context) error
	RemoveItemFromCart(c echo.Context) error
}

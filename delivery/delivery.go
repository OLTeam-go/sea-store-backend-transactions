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
	CheckoutCart(c echo.Context) error
	AcceptTransaction(c echo.Context) error
	RejectTransaction(c echo.Context) error
	FetchTransactions(c echo.Context) error
}

package service

import (
	"github.com/BazaarTrade/OrderMatchingService/internal/models"
)

type Exchanger interface {
	AddOrderBook(symbol string) error
	DeleteOrderBook(symbol string) error

	PlaceOrder(order models.PlaceOrderReq) ([]models.Order, error)
	CancelOrder(orderID int64) (models.Order, error)

	GetOrders(userID int64) ([]models.Order, error)
	GetCurrentOrders(userID int64) ([]models.Order, error)
}

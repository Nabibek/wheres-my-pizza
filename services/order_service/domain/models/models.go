// order.go - Core domain entity for the order
package domain

import (
	"time"
)

type Order struct {
	ID             int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	OrderNumber    string
	CustomerName   string
	OrderType      string
	TableNumber    *int
	DeliveryAddress *string
	TotalAmount    float64
	Priority       int
	Status         string
	Items          []OrderItem
}

type OrderItem struct {
	ID       int
	OrderID  int
	Name     string
	Quantity int
	Price    float64
}

type OrderStatusLog struct {
	ID          int
	OrderID     int
	Status      string
	ChangedBy   string
	ChangedAt   time.Time
	Notes       string
}

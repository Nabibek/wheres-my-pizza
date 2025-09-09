// order.go - Core domain entity for the order
package models

import (
	"time"
)

type CreateOrderRequest struct {
    CustomerName   string        `json:"customer_name"`
    OrderType      string        `json:"order_type"`
    TableNumber    *int          `json:"table_number,omitempty"`
    DeliveryAddress *string      `json:"delivery_address,omitempty"`
    Items          []OrderItem   `json:"items"`
}


type OrderMessage struct {
    OrderNumber     string       `json:"order_number"`
    CustomerName    string       `json:"customer_name"`
    OrderType       string       `json:"order_type"`
    TableNumber     *int         `json:"table_number,omitempty"`
    DeliveryAddress *string      `json:"delivery_address,omitempty"`
    Items           []OrderItem  `json:"items"`
    TotalAmount     float64      `json:"total_amount"`
    Priority        int          `json:"priority"`
}


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

// order_service.go - The core logic for creating orders
package service

import (
	"fmt"
	"restaurant-system/services/order-service/domain/models"
	"restaurant-system/services/order-service/domain/ports"
	"time"
)

type OrderService struct {
	OrderRepository   ports.OrderRepository
	OrderItemRepo     ports.OrderItemRepository
	StatusLogRepo     ports.OrderStatusLogRepository
	RabbitMQPublisher ports.RabbitMQPublisher
}

func (s *OrderService) CreateOrder(customerName, orderType string, items []models.OrderItem, tableNumber *int, deliveryAddress *string) (models.Order, error) {
	// Validate order
	if err := validateOrder(customerName, orderType, items, tableNumber, deliveryAddress); err != nil {
		return models.Order{}, err
	}

	// Calculate total amount and priority
	totalAmount := calculateTotalAmount(items)
	priority := calculatePriority(totalAmount)

	// Generate order number
	// orderNumber := generateOrderNumber()

	// Start a transaction for database operations
	order := models.Order{
		CustomerName:    customerName,
		OrderType:       orderType,
		TableNumber:     tableNumber,
		DeliveryAddress: deliveryAddress,
		TotalAmount:     totalAmount,
		Priority:        priority,
		Status:          "received",
		Items:           items,
	}

	// Insert order into the database
	err := s.OrderRepository.SaveOrder(order)
	if err != nil {
		return models.Order{}, err
	}

	// Insert items into the order_items table
	for _, item := range items {
		item.OrderID = order.ID
		err := s.OrderItemRepo.SaveOrderItem(item)
		if err != nil {
			return models.Order{}, err
		}
	}

	// Log the status
	statusLog := models.OrderStatusLog{
		OrderID:   order.ID,
		Status:    "received",
		ChangedBy: "system",
		ChangedAt: time.Now(),
	}
	err = s.StatusLogRepo.SaveOrderStatusLog(statusLog)
	if err != nil {
		return models.Order{}, err
	}

	// Publish to RabbitMQ
	err = s.RabbitMQPublisher.PublishOrder(order)
	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

// Helper functions
func validateOrder(customerName, orderType string, items []models.OrderItem, tableNumber *int, deliveryAddress *string) error {
	// Add input validation logic here (length checks, etc.)
	return nil
}

func calculateTotalAmount(items []models.OrderItem) float64 {
	var total float64
	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}

func calculatePriority(totalAmount float64) int {
	switch {
	case totalAmount > 100:
		return 10
	case totalAmount >= 50 && totalAmount <= 100:
		return 5
	default:
		return 1
	}
}

func generateOrderNumber() string {
	// Generate order number based on the current date and an incremental counter
	return fmt.Sprintf("ORD_%s_%03d", time.Now().Format("20060102"), 1) // Just for example
}

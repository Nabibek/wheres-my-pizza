package service

import (
	"fmt"
	"regexp"
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
	orderNumber, err := s.generateOrderNumber()
	if err != nil {
		return models.Order{}, err
	}

	// Create order object
	order := models.Order{
		OrderNumber:     orderNumber,
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
	err = s.OrderRepository.SaveOrder(&order)
	if err != nil {
		return models.Order{}, fmt.Errorf("failed to save order: %w", err)
	}

	// Insert items into the order_items table
	for i := range items {
		items[i].OrderID = order.ID
		err := s.OrderItemRepo.SaveOrderItem(items[i])
		if err != nil {
			return models.Order{}, fmt.Errorf("failed to save order item: %w", err)
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
		return models.Order{}, fmt.Errorf("failed to save status log: %w", err)
	}

	// Publish to RabbitMQ
	err = s.RabbitMQPublisher.PublishOrder(order)
	if err != nil {
		return models.Order{}, fmt.Errorf("failed to publish order: %w", err)
	}

	return order, nil
}

// Enhanced validation function
func validateOrder(customerName, orderType string, items []models.OrderItem, tableNumber *int, deliveryAddress *string) error {
	// Validate customer name
	if customerName == "" {
		return fmt.Errorf("customer_name is required")
	}
	if len(customerName) > 100 {
		return fmt.Errorf("customer_name must be 100 characters or less")
	}

	// Validate customer name contains only allowed characters
	validNameRegex := regexp.MustCompile(`^[a-zA-Z\s\-']+$`)
	if !validNameRegex.MatchString(customerName) {
		return fmt.Errorf("customer_name contains invalid characters")
	}

	// Validate order type
	validOrderTypes := map[string]bool{
		"dine_in":  true,
		"takeout":  true,
		"delivery": true,
	}
	if !validOrderTypes[orderType] {
		return fmt.Errorf("order_type must be one of: dine_in, takeout, delivery")
	}

	// Validate items
	if len(items) == 0 {
		return fmt.Errorf("items cannot be empty")
	}
	if len(items) > 20 {
		return fmt.Errorf("maximum 20 items allowed per order")
	}

	for i, item := range items {
		if item.Name == "" {
			return fmt.Errorf("item[%d].name is required", i)
		}
		if len(item.Name) > 50 {
			return fmt.Errorf("item[%d].name must be 50 characters or less", i)
		}
		if item.Quantity < 1 || item.Quantity > 10 {
			return fmt.Errorf("item[%d].quantity must be between 1 and 10", i)
		}
		if item.Price < 0.01 || item.Price > 999.99 {
			return fmt.Errorf("item[%d].price must be between 0.01 and 999.99", i)
		}
	}

	// Conditional validation based on order type
	switch orderType {
	case "dine_in":
		if tableNumber == nil {
			return fmt.Errorf("table_number is required for dine_in orders")
		}
		if *tableNumber < 1 || *tableNumber > 100 {
			return fmt.Errorf("table_number must be between 1 and 100")
		}
		if deliveryAddress != nil {
			return fmt.Errorf("delivery_address should not be provided for dine_in orders")
		}
	case "delivery":
		if deliveryAddress == nil {
			return fmt.Errorf("delivery_address is required for delivery orders")
		}
		if len(*deliveryAddress) < 10 {
			return fmt.Errorf("delivery_address must be at least 10 characters")
		}
		if tableNumber != nil {
			return fmt.Errorf("table_number should not be provided for delivery orders")
		}
	case "takeout":
		if tableNumber != nil {
			return fmt.Errorf("table_number should not be provided for takeout orders")
		}
		if deliveryAddress != nil {
			return fmt.Errorf("delivery_address should not be provided for takeout orders")
		}
	}

	return nil
}

// Enhanced order number generation
func (s *OrderService) generateOrderNumber() (string, error) {
	datePrefix := time.Now().UTC().Format("20060102")

	// Use a simple timestamp-based approach for sequence number
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	sequence := timestamp % 1000

	return fmt.Sprintf("ORD_%s_%03d", datePrefix, sequence), nil
}

// Helper functions
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

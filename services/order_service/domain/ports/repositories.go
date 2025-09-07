// ports/repositories.go - Ports for database operations
package ports

import "restaurant-system/domain"

type OrderRepository interface {
	SaveOrder(order domain.Order) error
}

type OrderItemRepository interface {
	SaveOrderItem(item domain.OrderItem) error
}

type OrderStatusLogRepository interface {
	SaveOrderStatusLog(log domain.OrderStatusLog) error
}

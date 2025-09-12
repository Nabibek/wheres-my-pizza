package ports

import (
	"restaurant-system/services/order-service/domain/models"
)

type OrderRepository interface {
	SaveOrder(order models.Order) error
}

type OrderItemRepository interface {
	SaveOrderItem(item models.OrderItem) error
}

type OrderStatusLogRepository interface {
	SaveOrderStatusLog(log models.OrderStatusLog) error
}

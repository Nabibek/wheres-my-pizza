package ports

import (
	"restaurant-system/services/tracking-service/domain/models"
)

type OrderRepository interface {
	GetOrderByNumber(orderNumber string) (models.OrderStatusResponse, error)
	GetOrderStatusHistory(orderNumber string) ([]models.StatusHistory, error)
}

type WorkerRepository interface {
	GetAllWorkersStatus() ([]models.WorkerStatus, error)
}

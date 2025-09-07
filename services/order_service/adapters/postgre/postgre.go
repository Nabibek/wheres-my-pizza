// adapters/postgre/order_repository.go
package postgre

import (
	"restaurant-system/domain"
	"restaurant-system/ports"
	"gorm.io/gorm"
)

type PostgresOrderRepository struct {
	DB *gorm.DB
}

func (r *PostgresOrderRepository) SaveOrder(order domain.Order) error {
	return r.DB.Create(&order).Error
}

type PostgresOrderItemRepository struct {
	DB *gorm.DB
}

func (r *PostgresOrderItemRepository) SaveOrderItem(item domain.OrderItem) error {
	return r.DB.Create(&item).Error
}

type PostgresOrderStatusLogRepository struct {
	DB *gorm.DB
}

func (r *PostgresOrderStatusLogRepository) SaveOrderStatusLog(log domain.OrderStatusLog) error {
	return r.DB.Create(&log).Error
}

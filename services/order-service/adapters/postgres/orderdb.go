package postgres

import (
	"context"
	"restaurant-system/services/order-service/domain/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresOrderRepository struct {
	DB *pgxpool.Pool
}

func (r *PostgresOrderRepository) SaveOrder(order *models.Order) error {
	query := `
		INSERT INTO orders (number, customer_name, type, table_number, delivery_address, total_amount, priority, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at, updated_at
	`

	return r.DB.QueryRow(context.Background(), query,
		order.OrderNumber,
		order.CustomerName,
		order.OrderType,
		order.TableNumber,
		order.DeliveryAddress,
		order.TotalAmount,
		order.Priority,
		order.Status,
	).Scan(&order.ID, &order.CreatedAt, &order.UpdatedAt)
}

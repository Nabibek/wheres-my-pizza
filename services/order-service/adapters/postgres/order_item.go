package postgres

import (
	"context"
	"restaurant-system/services/order-service/domain/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresOrderItemRepository struct {
	DB *pgxpool.Pool
}

func (r *PostgresOrderItemRepository) SaveOrderItem(item models.OrderItem) error {
	query := `
		INSERT INTO order_items (order_id, name, quantity, price)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`
	
	return r.DB.QueryRow(context.Background(), query,
		item.OrderID,
		item.Name,
		item.Quantity,
		item.Price,
	).Scan(&item.ID, &item.CreatedAt)
}
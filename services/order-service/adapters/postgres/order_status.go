package postgres

import (
	"context"
	"restaurant-system/services/order-service/domain/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresOrderStatusLogRepository struct {
	DB *pgxpool.Pool
}

func (r *PostgresOrderStatusLogRepository) SaveOrderStatusLog(log models.OrderStatusLog) error {
	query := `
		INSERT INTO order_status_log (order_id, status, changed_by, changed_at, notes)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`
	
	return r.DB.QueryRow(context.Background(), query,
		log.OrderID,
		log.Status,
		log.ChangedBy,
		log.ChangedAt,
		log.Notes,
	).Scan(&log.ID, &log.CreatedAt)
}
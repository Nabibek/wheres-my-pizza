package postgres

import (
	"context"
	"restaurant-system/services/tracking-service/domain/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresOrderRepository struct {
	DB *pgxpool.Pool
}

func (r *PostgresOrderRepository) GetOrderByNumber(orderNumber string) (models.OrderStatusResponse, error) {
	query := `
		SELECT 
			number, 
			status, 
			updated_at, 
			completed_at as estimated_completion, 
			processed_by
		FROM orders 
		WHERE number = $1
	`

	var statusResponse models.OrderStatusResponse
	err := r.DB.QueryRow(context.Background(), query, orderNumber).Scan(
		&statusResponse.OrderNumber,
		&statusResponse.CurrentStatus,
		&statusResponse.UpdatedAt,
		&statusResponse.EstimatedCompletion,
		&statusResponse.ProcessedBy,
	)

	if err != nil {
		return models.OrderStatusResponse{}, err
	}

	return statusResponse, nil
}

func (r *PostgresOrderRepository) GetOrderStatusHistory(orderNumber string) ([]models.StatusHistory, error) {
	query := `
		SELECT 
			osl.status, 
			osl.changed_at as timestamp, 
			osl.changed_by
		FROM order_status_log osl
		JOIN orders o ON osl.order_id = o.id
		WHERE o.number = $1
		ORDER BY osl.changed_at ASC
	`

	rows, err := r.DB.Query(context.Background(), query, orderNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []models.StatusHistory
	for rows.Next() {
		var entry models.StatusHistory
		err := rows.Scan(
			&entry.Status,
			&entry.Timestamp,
			&entry.ChangedBy,
		)
		if err != nil {
			return nil, err
		}
		history = append(history, entry)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return history, nil
}
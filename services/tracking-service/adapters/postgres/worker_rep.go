package postgres

import (
	"context"
	"restaurant-system/services/tracking-service/domain/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresWorkerRepository struct {
	DB *pgxpool.Pool
}

func (r *PostgresWorkerRepository) GetAllWorkersStatus() ([]models.WorkerStatus, error) {
	query := `
		SELECT 
			name as worker_name, 
			orders_processed, 
			last_seen
		FROM workers
		ORDER BY name
	`

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workers []models.WorkerStatus
	for rows.Next() {
		var worker models.WorkerStatus
		err := rows.Scan(
			&worker.WorkerName,
			&worker.OrdersProcessed,
			&worker.LastSeen,
		)
		if err != nil {
			return nil, err
		}
		workers = append(workers, worker)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return workers, nil
}
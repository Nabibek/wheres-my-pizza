package service

import (
	"log"
	"restaurant-system/services/tracking-service/domain/models"
	"restaurant-system/services/tracking-service/domain/ports"
	"time"
)

type TrackingService struct {
	OrderRepo   ports.OrderRepository
	WorkerRepo  ports.WorkerRepository
}

func NewTrackingService(orderRepo ports.OrderRepository, workerRepo ports.WorkerRepository) *TrackingService {
	return &TrackingService{
		OrderRepo:  orderRepo,
		WorkerRepo: workerRepo,
	}
}

func (s *TrackingService) GetOrderStatus(orderNumber string) (models.OrderStatusResponse, error) {
	log.Printf("Getting status for order: %s", orderNumber)
	return s.OrderRepo.GetOrderByNumber(orderNumber)
}

func (s *TrackingService) GetOrderHistory(orderNumber string) ([]models.StatusHistory, error) {
	log.Printf("Getting history for order: %s", orderNumber)
	return s.OrderRepo.GetOrderStatusHistory(orderNumber)
}

func (s *TrackingService) GetWorkersStatus() ([]models.WorkerStatus, error) {
	log.Printf("Getting status for all workers")
	workers, err := s.WorkerRepo.GetAllWorkersStatus()
	if err != nil {
		return nil, err
	}

	// Determine online/offline status based on last_seen
	for i := range workers {
		if time.Since(workers[i].LastSeen) > 2*time.Minute { // 2 minutes threshold
			workers[i].Status = "offline"
		} else {
			workers[i].Status = "online"
		}
	}

	return workers, nil
}
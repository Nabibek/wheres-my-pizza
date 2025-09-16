package web

import (
	"encoding/json"
	"log"
	"net/http"
	"restaurant-system/services/order-service/domain/models"
	"restaurant-system/services/order-service/domain/service"
	"strings"
)

type WebHandler struct {
	OrderService service.OrderService
}

func NewWebHandler(OrderService service.OrderService) *WebHandler {
	return &WebHandler{OrderService: OrderService}
}

func (h *WebHandler) HandleOrder(w http.ResponseWriter, r *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")

	var request struct {
		CustomerName    string             `json:"customer_name"`
		OrderType       string             `json:"order_type"`
		Items           []models.OrderItem `json:"items"`
		TableNumber     *int               `json:"table_number,omitempty"`
		DeliveryAddress *string            `json:"delivery_address,omitempty"`
	}

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("Invalid request: %v", err)
		http.Error(w, `{"error": "Invalid JSON format"}`, http.StatusBadRequest)
		return
	}

	// Create order
	order, err := h.OrderService.CreateOrder(
		request.CustomerName,
		request.OrderType,
		request.Items,
		request.TableNumber,
		request.DeliveryAddress,
	)
	if err != nil {
		log.Printf("Failed to create order: %v", err)
		// Check if it's a validation error
		if strings.Contains(err.Error(), "validation") {
			http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusBadRequest)
		} else {
			http.Error(w, `{"error": "Failed to create order"}`, http.StatusInternalServerError)
		}
		return
	}

	// Respond with the created order
	response := map[string]interface{}{
		"order_number": order.OrderNumber,
		"status":       order.Status,
		"total_amount": order.TotalAmount,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

package web

type WebHandler struct{}

func NewWebHandler() *WebHandler {
	return &WebHandler{}
}

func (h *WebHandler) HandleOrder(w http.ResponseWriter, r *http.Request) {
	var request struct {
		CustomerName   string          `json:"customer_name"`
		OrderType      string          `json:"order_type"`
		Items          []domain.OrderItem `json:"items"`
		TableNumber    *int            `json:"table_number,omitempty"`
		DeliveryAddress *string        `json:"delivery_address,omitempty"`
	}

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Create order
	order, err := h.OrderService.CreateOrder(request.CustomerName, request.OrderType, request.Items, request.TableNumber, request.DeliveryAddress)
	if err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	// Respond with the created order
	response := map[string]interface{}{
		"order_number": order.OrderNumber,
		"status":       order.Status,
		 "total_amount": order.TotalAmount,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

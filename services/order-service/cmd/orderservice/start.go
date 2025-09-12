package orderservice

import (
	// "fmt"
	"log"
	// "restaurant-system/services/order-service/domain/service"
	"restaurant-system/services/order-service/adapters/rabbitmq"
	// "restaurant-system/services/order-service/adapters/postgres"
)

func OrderService() {
	// Connect to RabbitMQ
	client, err := rabbitmq.NewClient("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	defer client.Close()

	err = client.DeclareExchange("orders_topic")
	if err != nil {
		log.Fatal("Failed to declare exchange:", err)
	}

	message := `{"order_number": "ORD_001", "customer_name": "Alice"}`
	err = client.Publish("orders_topic", "kitchen.takeout.1", message)
	if err != nil {
		log.Fatal("Failed to publish:", err)
	}

	// // Set up the repositories and services
	// orderRepo := &postgres.PostgresOrderRepository{DB: db}
	// orderItemRepo := &postgres.PostgresOrderItemRepository{DB: db}
	// orderStatusLogRepo := &postgres.PostgresOrderStatusLogRepository{DB: db}
	// rabbitMQPublisher := &rabbitmq.RabbitMQPublisher{
	// 	Connection: conn,
	// 	Channel:    channel,
	// }

	// orderService := &service.OrderService{
	// 	OrderRepository:   orderRepo,
	// 	OrderItemRepo:     orderItemRepo,
	// 	StatusLogRepo:     orderStatusLogRepo,
	// 	RabbitMQPublisher: rabbitMQPublisher,
	// }

}

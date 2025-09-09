package orderservice

import (
	"log"
	"net/http"
	"restaurant-system/adapters/http"
	"restaurant-system/adapters/postgres"
	"restaurant-system/adapters/rabbitmq"
	"restaurant-system/service"

	"github.com/streadway/amqp"
)

func OrderService() {
	// Connect to PostgreSQL
	db, err := gorm.Open(postgres.Open("your_database_url"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	channel, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to create RabbitMQ channel:", err)
	}

	// Set up the repositories and services
	orderRepo := &postgres.PostgresOrderRepository{DB: db}
	orderItemRepo := &postgres.PostgresOrderItemRepository{DB: db}
	orderStatusLogRepo := &postgres.PostgresOrderStatusLogRepository{DB: db}
	rabbitMQPublisher := &rabbitmq.RabbitMQPublisher{
		Connection: conn,
		Channel:    channel,
	}

	orderService := &service.OrderService{
		OrderRepository:   orderRepo,
		OrderItemRepo:     orderItemRepo,
		StatusLogRepo:     orderStatusLogRepo,
		RabbitMQPublisher: rabbitMQPublisher,
	}

	// Set up HTTP handler
	orderHandler := &http.OrderHandler{
		OrderService: orderService,
	}

	http.HandleFunc("/orders", orderHandler.CreateOrder)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

// func main() {
//     url := os.Getenv("RABBITMQ_URL")
//     if url == "" {
//         url = "amqp://guest:guest@localhost:5672/"
//     }

//     client, err := rabbitmq.NewClient(url)
//     if err != nil {
//         log.Fatal("Failed to connect to RabbitMQ:", err)
//     }
//     defer client.Close()

//     log.Println("Order service started")

//     // Example: publish a test order
//     err = client.PublishOrder([]byte(`{"order":"pizza"}`), "kitchen.takeout.1")
//     if err != nil {
//         log.Fatal("Failed to publish order:", err)
//     }
// }

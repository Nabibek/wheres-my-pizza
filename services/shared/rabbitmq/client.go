package rabbitmq

import (
    "log"
    "github.com/rabbitmq/amqp091-go"
    "os"
    "sync"
)

type RabbitMQClient struct {
    Connection *amqp091.Connection
    Channel    *amqp091.Channel
}

var (
    once      sync.Once          // ensures the connection is created once
    client    *RabbitMQClient    // singleton RabbitMQ client
    connMutex sync.Mutex         // mutex for thread-safety
)

// NewClient creates a RabbitMQ client, but only once for the entire application.
func NewClient() (*RabbitMQClient, error) {
    // Ensure the connection is created only once
    once.Do(func() {
        var err error

        rabbitmqURL := os.Getenv("RABBITMQ_URL")
        if rabbitmqURL == "" {
            rabbitmqURL = "amqp://guest:guest@localhost:5672/"
        }

        client, err = initializeConnection(rabbitmqURL)
        if err != nil {
            log.Fatalf("Failed to connect to RabbitMQ: %v", err)
        }
    })

    return client, nil
}

// initializeConnection establishes the connection and initializes the channel.
func initializeConnection(rabbitmqURL string) (*RabbitMQClient, error) {
    conn, err := amqp091.Dial(rabbitmqURL)
    if err != nil {
        return nil, err
    }

    ch, err := conn.Channel()
    if err != nil {
        return nil, err
    }

    log.Println("Successfully connected to RabbitMQ")
    return &RabbitMQClient{
        Connection: conn,
        Channel:    ch,
    }, nil
}

// Publish a message to a specific exchange and routing key
func (r *RabbitMQClient) Publish(exchange, routingKey, body string) error {
    return r.Channel.Publish(
        exchange,    // exchange
        routingKey,  // routing key
        false,        // mandatory
        false,        // immediate
        amqp091.Publishing{
            ContentType: "text/plain",
            Body:        []byte(body),
        },
    )
}

// Consume messages from a queue
func (r *RabbitMQClient) Consume(queueName string) (<-chan amqp091.Delivery, error) {
    return r.Channel.Consume(
        queueName, // queue name
        "",        // consumer tag
        true,      // auto-ack
        false,     // exclusive
        false,     // no-local
        false,     // no-wait
        nil,       // arguments
    )
}

// Close the connection and channel gracefully
func (r *RabbitMQClient) Close() {
    connMutex.Lock()
    defer connMutex.Unlock()

    if err := r.Channel.Close(); err != nil {
        log.Printf("Failed to close channel: %v", err)
    }
    if err := r.Connection.Close(); err != nil {
        log.Printf("Failed to close connection: %v", err)
    }
}

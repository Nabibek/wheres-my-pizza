// internal/rabbitmq/publisher.go
package rabbitmq

import (
    "context"
    "log"
    "time"

    "github.com/rabbitmq/amqp091-go"
)

func (c *Client) PublishOrder(order []byte, routingKey string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    return c.Channel.PublishWithContext(ctx,
        "orders_topic", // exchange
        routingKey,     // routing key like kitchen.takeout.1
        false,
        false,
        amqp091.Publishing{
            ContentType: "application/json",
            DeliveryMode: amqp091.Persistent,
            Body: order,
        },
    )
}

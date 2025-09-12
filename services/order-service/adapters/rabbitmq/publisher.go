// internal/rabbitmq/publisher.go
package rabbitmq

import (
	"context"
	"time"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (c *Client) Publish(exchange, routingKey, body string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.channel.PublishWithContext(
		ctx,
		exchange,
		routingKey,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         []byte(body),
			DeliveryMode: amqp.Persistent, // persist on disk
		},
	)
}

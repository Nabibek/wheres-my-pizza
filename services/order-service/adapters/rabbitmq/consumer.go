// internal/rabbitmq/consumer.go
package rabbitmq

import (
	"log"
)

func (c *Client) ConsumeOrders(queue string) {
	msgs, err := c.Channel.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to consume messages:", err)
	}

	for msg := range msgs {
		log.Printf("Kitchen received order: %s", msg.Body)
		// TODO: process order (call service layer)
	}
}

// internal/rabbitmq/consumer.go
package rabbitmq

func (c *Client) Consume(queue string, handler func(string)) error {
	msgs, err := c.channel.Consume(
		queue,
		"",    // consumer
		false, // auto-ack (false → manual ack)
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			handler(string(msg.Body))
			msg.Ack(false)
		}
	}()
	return nil
}

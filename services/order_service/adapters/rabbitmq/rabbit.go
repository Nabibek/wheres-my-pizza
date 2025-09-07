// adapters/rabbitmq/rabbitmq_publisher.go
package rabbit

import (
	"encoding/json"
	"restaurant-system/domain"
	"restaurant-system/ports"
	"github.com/streadway/amqp"
)

type RabbitMQPublisher struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func (r *RabbitMQPublisher) PublishOrder(order domain.Order) error {
	message := map[string]interface{}{
		"order_number":    order.OrderNumber,
		"customer_name":   order.CustomerName,
		"order_type":      order.OrderType,
		"table_number":    order.TableNumber,
		"delivery_address": order.DeliveryAddress,
		"items":           order.Items,
		"total_amount":    order.TotalAmount,
		"priority":        order.Priority,
	}

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	// Create a persistent message
	err = r.Channel.Publish(
		"orders_topic",               // Exchange
		"kitchen." + order.OrderType + "." + string(order.Priority), // Routing key
		true,                         // Mandatory
		false,                        // Immediate
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         body,
		},
	)
	return err
}

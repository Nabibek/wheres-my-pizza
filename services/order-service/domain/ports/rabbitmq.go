// ports/rabbitmq.go - Port for RabbitMQ Publisher
package ports

import "restaurant-system/domain"

type RabbitMQPublisher interface {
	PublishOrder(order domain.Order) error
}

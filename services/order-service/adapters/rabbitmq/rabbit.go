// internal/rabbitmq/client.go
package rabbitmq

import (
    "log"
    "github.com/rabbitmq/amqp091-go"
)

type Client struct {
    Conn    *amqp091.Connection
    Channel *amqp091.Channel
}

func NewClient(url string) (*Client, error) {
    conn, err := amqp091.Dial(url)
    if err != nil {
        return nil, err
    }

    ch, err := conn.Channel()
    if err != nil {
        return nil, err
    }

    return &Client{Conn: conn, Channel: ch}, nil
}

func ExchangeDeclare(ch *amqp091.Channel){
    err = ch.ExchangeDeclare(
        "orders_topic", // name
        "topic",        // type
        true,           // durable
        false,          // auto-delete
        false,          // internal
        false,          // no-wait
        nil,            // args
    )
    if err != nil {
        log.Fatal("Failed to declare exchange:", err)
    }
}

func QueueDeclare(ch *amqp091.Channel){
    q, err := ch.QueueDeclare(
        "kitchen_takeout", // queue name
        true,              // durable
        false,             // delete when unused
        false,             // exclusive
        false,             // no-wait
        nil,               // arguments
    )
    if err != nil {
        log.Fatal("Failed to declare queue:", err)
    }
}


func (c *Client) Close() {
    c.Channel.Close()
    c.Conn.Close()
}

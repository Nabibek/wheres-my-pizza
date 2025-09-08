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

func (c *Client) Close() {
    c.Channel.Close()
    c.Conn.Close()
}

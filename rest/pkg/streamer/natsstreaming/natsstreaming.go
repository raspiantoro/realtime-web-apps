package natsstreaming

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	stan "github.com/nats-io/stan.go"
)

// Client define client implementation for nats streaming
type Client struct {
	conn stan.Conn
}

//ClientOption natsstreaming client options
type ClientOption struct {
	ClusterID string
	ClientID  string
	Host      string
	Port      uint64
}

// Callback callback function when message received
type Callback func(data []byte, ack Ack) error

// Message callback return message
type Message struct {
	data []byte
	ack  Ack
}

// Ack callback return ack handler
type Ack func() error

// ListenerValue listener return value
type ListenerValue struct {
	Unsubscribe func() error
}

//NewNatsStreamingClient create nats streaming client instance
func NewNatsStreamingClient(opt ClientOption) (c *Client, err error) {
	c = new(Client)

	url := c.getURL(opt.Host, opt.Port)

	c.conn, err = stan.Connect(
		opt.ClusterID,
		opt.ClientID,
		stan.NatsURL(url),
	)

	if err != nil {
		err = fmt.Errorf("stan error connect to %s: %s", url, err)
	}
	return
}

func (c *Client) getURL(host string, port uint64) (url string) {

	url = stan.DefaultNatsURL

	if host != "" {
		url = strings.ReplaceAll(url, "127.0.0.1", host)
	}

	if port != 0 {
		url = strings.ReplaceAll(url, "4222", strconv.FormatUint(port, 10))
	}

	return
}

// Emit emit/publish to any topic on nats streaming cluster
func (c *Client) Emit(topic string, payload []byte) (err error) {
	err = c.conn.Publish(topic, payload)
	return
}

// Listen listen/publish to any topic on nats streaming cluster
func (c *Client) Listen(topic string, durable string, handler Callback) (listenerValue ListenerValue, err error) {
	aw, err := time.ParseDuration("60s")
	if err != nil {
		return
	}

	sub, err := c.conn.Subscribe(topic, func(msg *stan.Msg) {
		handler(msg.Data, msg.Ack)
	},
		stan.DurableName(durable),
		stan.SetManualAckMode(),
		stan.AckWait(aw),
	)
	if err != nil {
		return
	}

	listenerValue.Unsubscribe = sub.Unsubscribe

	return
}

//Close close nats streaming connection
func (c *Client) Close() (err error) {
	err = c.conn.Close()
	return
}

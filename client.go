// Package msgr is the latest iteration of the message queueing package.
package msgr

import (
	"log"
	"strings"
	"sync"

	"github.com/streadway/amqp"
)

type (
	// Client defines the base message queueing behavior
	Client interface {
		Dial() error
		Close()
	}
	// Producer defines the base message producing behavior
	Producer interface {
		Post([]byte) bool
		Close()
	}
	// Consumer defines the base message consuming behavior
	Consumer interface {
		Accept() (bool, <-chan amqp.Delivery)
		Close()
	}
)

type (
	// Config is the queue configuration settings.
	Config struct {
		URI     string
		Channel string
	}
	// QueueProducer implements Producer.
	QueueProducer struct {
		*QueueClient
	}
	// QueueConsumer implements Consumer.
	QueueConsumer struct {
		*QueueClient
	}
	// QueueClient implements the client behavior.
	QueueClient struct {
		conn    *amqp.Connection
		channel *amqp.Channel
		mu      sync.Mutex
		conf    *Config
	}
)

// ConnectP returns a producer.
func ConnectP(conf *Config) *QueueProducer {
	c := &QueueClient{
		conf: conf,
	}
	c.Dial()
	return &QueueProducer{c}
}

// ConnectC returns a consumer.
func ConnectC(conf *Config) *QueueConsumer {
	c := &QueueClient{
		conf: conf,
	}
	c.Dial()
	return &QueueConsumer{c}
}

// Dial makes a connection.
func (c *QueueClient) Dial() {
	c.conn = dial(c.conf.URI)
}

// Close closes a connection.
func (c *QueueClient) Close() {
	err := c.conn.Close()
	if err != nil {
		// Ignore 504 - channel/connection is not open
		if !strings.Contains(err.Error(), "Exception (504)") {
			log.Println(err)
		}
	}
}

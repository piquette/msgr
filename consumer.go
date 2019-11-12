package msgr

import (
	"log"

	"github.com/streadway/amqp"
)

// Accept delivers a stream of messgaes.
func (c *QueueConsumer) Accept() (bool, <-chan amqp.Delivery) {
	// check channel state.
	c.mu.Lock()
	if c.channel == nil {
		newchan := channel(c.conn, c.conf.Channel)
		if newchan == nil {
			c.mu.Unlock()
			return false, nil
		}
		c.channel = newchan
	}
	c.mu.Unlock()

	messages, err := c.channel.Consume(
		c.conf.Channel,   // queue
		"",               // messageConsumer
		consumeAutoAck,   // auto-ack
		consumeExclusive, // exclusive
		false,            // no-local
		consumeNoWait,    // no-wait
		nil,              // args
	)
	if err != nil {
		log.Println("could not consume messages")
		log.Println("error was: ", err)
		c.mu.Lock()
		c.channel = nil
		c.mu.Unlock()
		return false, nil
	}
	return true, messages
}

package msgr

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

var (
	// Internal default settings.
	// WARNING- these should only be modified during tests.
	// Producer defaults.
	publishMandatory = true  // mandatory requires a queue to already be declared.
	publishImmediate = false // immediate requires a consumer to be available on the other end.
	// Queue defaults.
	queueDurable   = true  // durable will survive server restarts
	queueDelete    = false // delete will close a queue when there are no consumers or bindings
	queueExclusive = false // exclusive queues are only accessible by the connection that declares them
	queueNoWait    = false // no-wait will assume the queue was declared on the server
	// Consumer defaults.
	consumeAutoAck   = false // auto-ack is false, the consumer should always call Delivery.Ack
	consumeExclusive = false // exclusive ensures that this is the sole consumer of a queue
	consumeNoWait    = false // do not wait for the server to confirm the request and immediately begin deliveries
)

func dial(uri string) *amqp.Connection {
	for {
		conn, err := amqp.Dial(uri)
		if err == nil {
			log.Println("connected to queue")
			return conn
		}

		log.Println("dial failed. retrying in 1s: ", err)
		time.Sleep(1000 * time.Millisecond)
	}
}

func channel(conn *amqp.Connection, name string) *amqp.Channel {
	for {
		channel, err := conn.Channel()
		if err == nil {

			_, err = channel.QueueDeclare(
				name,           // name
				queueDurable,   // durable
				queueDelete,    // delete when unused
				queueExclusive, // exclusive
				queueNoWait,    // no-wait
				nil,            // arguments
			)
			if err == nil {
				return channel
			}
		}

		if err == amqp.ErrClosed {
			return nil
		}

		log.Println("channel open failed. retrying in 1s: ", err)
		time.Sleep(1000 * time.Millisecond)
	}
}

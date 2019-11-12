package msgr

import (
	"log"

	"github.com/streadway/amqp"
)

// Post sends a message.
func (p *QueueProducer) Post(msg []byte) bool {
	// check channel state.
	p.mu.Lock()
	if p.channel == nil {
		newchan := channel(p.conn, p.conf.Channel)
		if newchan == nil {
			p.mu.Unlock()
			return false
		}
		p.channel = newchan
	}
	p.mu.Unlock()

	// post.
	err := p.post(msg)
	if err != nil {
		log.Println("could not publish message")
		log.Println("error was: ", err)
		p.mu.Lock()
		p.channel = nil
		p.mu.Unlock()
		return false
	}
	return true
}

func (p *QueueProducer) post(msg []byte) error {
	return p.channel.Publish(
		"",               // exchange
		p.conf.Channel,   // routing key
		publishMandatory, // mandatory
		publishImmediate, // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         msg,
		})
}

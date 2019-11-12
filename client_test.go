package msgr

import (
	"testing"

	tests "github.com/piquette/msgr/testing"
	"github.com/stretchr/testify/assert"
)

// Integration tests.

// TestQueueSuccess will smoke test the pub/sub capabilities of the queue client.
func TestQueueSuccess(t *testing.T) {
	done := make(chan bool, 1)
	testID := "1-1-1"
	conf := &Config{
		URI:     tests.AMQPServiceURL,
		Channel: tests.AMQPIntegrationQueue,
	}

	go func() {
		// Consumer.
		consumer := ConnectC(conf)
		open, messages := consumer.Accept()
		assert.True(t, open)
		for m := range messages {
			e := &tests.QueueEntry{}
			err := e.Decode(m.Body)
			assert.Nil(t, err)
			m.Ack(false)
			assert.Equal(t, e.ID, testID)
			// Close.
			consumer.Close()
		}
		done <- true
	}()

	qe := &tests.QueueEntry{
		ID: testID,
	}
	message, err := qe.Encode()
	assert.Nil(t, err)
	// Outbox.
	producer := ConnectP(conf)

	success := producer.Post(message)
	assert.True(t, success)

	// Close.
	producer.Close()

	<-done
	tests.Delete(tests.AMQPIntegrationQueue)
}

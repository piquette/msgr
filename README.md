# msgr | AMQP Client

[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/piquette/msgr) [![Build Status](https://travis-ci.org/piquette/msgr.svg?branch=master)](https://travis-ci.org/piquette/msgr) [![Coverage Status](https://coveralls.io/repos/github/piquette/msgr/badge.svg?branch=master)](https://coveralls.io/github/piquette/msgr?branch=master)

## Summary

This go module is a tiny wrapper around the [streadway/amqp](https://github.com/streadway/amqp), the standard low level golang client for [rabbitmq](https://www.rabbitmq.com/), the open source message broker. 

`msgr` abstracts away a few of the complexities of rabbitmq and exposes a slightly friendlier interface, intended for simple applications.

## Documentation

For details on all the functionality in this library, see the [GoDoc][godoc] documentation.

## Installation

This project supports modules and Go 1.13+. Add `msgr` to your own project the usual way -

```sh
go get github.com/piquette/msgr
```

## Usage example

### Producer

```go
// Instantiate and connect to the server.
conf := &msgr.Config{
    URI:     "amqp://localhost:5672",
    Channel: "queue_name",
}
producer = msgr.ConnectP(conf)
defer producer.Close()

// Send a message.
Enqueue:
{
    success := producer.Post([]byte("hi")])
    if !success {
        // Retry for all eternity.
        log.Println("could not enqueue msg")
        time.Sleep(time.Second * 3)
        goto Enqueue
    }
}
```

### Consumer

```go
// Instantiate and connect to the server.
conf := &msgr.Config{
    URI:     "amqp://localhost:5672",
    Channel: "queue_name",
}
consumer = msgr.ConnectC(conf)
defer consumer.Close()

// Receive messages.
open, messages := s.Consumer.Accept()
if !open {
    return
}
// Range over the messages chan.
for recv := range messages {
    // Got one.
    fmt.Println(string(recv.Body)) // prints 'hi'.

    // Don't forget to acknowledge.
    recv.Ack(false)
}
```


## Contributing

This modules is a work in progress and needs a lot of refinement. Please [submit an issue][issues] if you need help! 


[godoc]: http://godoc.org/github.com/piquette/msgr
[issues]: https://github.com/piquette/msgr/issues/new
[pulls]: https://github.com/piquette/msgr/pulls
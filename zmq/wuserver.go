package main

import (
	"fmt"
	"math/rand"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	publisher, _ := zmq.NewSocket(zmq.PUB)
	defer publisher.Close()

	// Bind server.
	publisher.Bind("tcp://*:5556")

	// Initialize random number generator.
	rand.Seed(time.Now().UnixNano())

	for {
		r := rand.Intn(5)

		msg := fmt.Sprintf("%d from PUB", r)
		publisher.Send(msg, 0)

		time.Sleep(2 * time.Second)
	}
}

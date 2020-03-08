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

	publisher.Connect("tcp://localhost:5556")

	for {
		r := rand.Intn(100)

		msg := fmt.Sprintf("Pub %d", r)
		publisher.Send(msg, 0)

		time.Sleep(2 * time.Second)
	}
}

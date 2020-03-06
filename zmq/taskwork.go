package main

// Work

import (
	"fmt"
	"math/rand"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	receiver, _ := zmq.NewSocket(zmq.PULL)
	defer receiver.Close()
	receiver.Connect("tcp://localhost:5557")

	sender, _ := zmq.NewSocket(zmq.PUSH)
	defer sender.Close()
	sender.Connect("tcp://localhost:5558")

	// initialize the default Source to a deterministic state.
	rand.Seed(time.Now().UnixNano())

	for {
		msg, _ := receiver.Recv(0)
		fmt.Printf("%v ", msg)

		// Make the worker busy.
		msec := rand.Intn(100) + 1
		time.Sleep(time.Duration(msec) * time.Millisecond)

		sender.Send("", 0)
	}
}

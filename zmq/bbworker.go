package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	receiver, _ := zmq.NewSocket(zmq.PULL)
	defer receiver.Close()
	receiver.Connect("tcp://localhost:5557")

	for {
		msg, _ := receiver.Recv(0)
		fmt.Println("Worker", msg)
	}
}

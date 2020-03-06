package main

// Sink

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	receiver, _ := zmq.NewSocket(zmq.PULL)
	defer receiver.Close()
	receiver.Bind("tcp://*:5558")

	fmt.Println("wait...")
	receiver.Recv(0)

	for i := 0; i < 100; i++ {
		receiver.Recv(0)
		if i%10 == 0 {
			fmt.Print(" ")
		}
		fmt.Print(".")
	}
}

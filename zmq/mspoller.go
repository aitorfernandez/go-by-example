package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	receiver, _ := zmq.NewSocket(zmq.PULL)
	defer receiver.Close()
	receiver.Connect("tcp://localhost:5557")

	subscriber, _ := zmq.NewSocket(zmq.SUB)
	defer subscriber.Close()
	subscriber.Connect("tcp://localhost:5556")
	subscriber.SetSubscribe("1")

	poller := zmq.NewPoller()
	// http://api.zeromq.org/4-1:zmq-getsockopt#toc8
	// Indicates that at least one message may be received from the specified socket without blocking.
	poller.Add(receiver, zmq.POLLIN)
	poller.Add(subscriber, zmq.POLLIN)

	//  Process messages from both sockets
	for {
		// If timeout < 0, wait forever until a matching event is detected
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {
			switch s := socket.Socket; s {
			case receiver:
				task, _ := s.Recv(0)
				fmt.Println("From vent", task)
			case subscriber:
				update, _ := s.Recv(0)
				fmt.Println("From sub", update)
			}
		}
	}
}

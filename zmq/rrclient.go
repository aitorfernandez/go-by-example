package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	client, _ := zmq.NewSocket(zmq.REQ)
	defer client.Close()
	client.Connect("tcp://localhost:5550")

	for i := 0; i < 10; i++ {
		client.Send(fmt.Sprintf("PING %d", i), 0)

		reply, _ := client.Recv(0)
		fmt.Println("Client reply %d, %s", i, reply)
	}
}

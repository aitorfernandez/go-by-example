package main

import (
	"fmt"
	"strconv"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	// Create socket
	client, _ := zmq.NewSocket(zmq.REQ)
	defer client.Close()

	client.Connect("tcp://localhost:5005")

	// The client issues zmq_send() and then zmq_recv(), in a loop (or once if that's all it needs). Doing any other sequence (e.g., sending two messages in a row) will result in a return code of -1
	for r := 1; r != 5; r++ {
		// The client sends a request and reads the reply back from the server.
		// zmq_send
		fmt.Println("Client send", r)
		client.Send(strconv.Itoa(r), 0)

		// and wait for the reply and reads with zmq_recv.
		reply, _ := client.Recv(0)
		fmt.Println("Client received ", reply)
	}
}

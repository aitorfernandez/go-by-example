package main

import (
	"strconv"

	zmq "github.com/pebbe/zmq4"

	"fmt"
	"time"
)

func main() {
	server, _ := zmq.NewSocket(zmq.REP)
	defer server.Close()

	// Bind to the port.
	server.Bind("tcp://*:5005")

	// And Binding here forever
	for {
		msg, _ := server.Recv(0)
		fmt.Println("Server received", msg)

		time.Sleep(2 * time.Second)

		// Update message and return.
		res, _ := strconv.Atoi(msg)

		fmt.Println("Server send", res)
		server.Send(strconv.Itoa(res*101), 0)
	}
}

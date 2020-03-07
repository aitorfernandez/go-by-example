package main

import (
	"fmt"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	responser, _ := zmq.NewSocket(zmq.REP)
	defer responser.Close()
	responser.Connect("tcp://localhost:5560")

	for {
		// wait for client request
		request, _ := responser.Recv(0)
		fmt.Println("Worker request %s", request)

		time.Sleep(2 * time.Second)

		responser.Send("PONG", 0)
	}
}

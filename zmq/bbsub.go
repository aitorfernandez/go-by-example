package main

import (
	"fmt"
	"strings"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	defer subscriber.Close()

	subscriber.Bind("tcp://*:5556")
	subscriber.SetSubscribe("myTopic")

	sender, _ := zmq.NewSocket(zmq.PUSH)
	defer sender.Close()
	sender.Bind("tcp://*:5557")

	for {
		msg, _ := subscriber.Recv(0)

		if msgs := strings.Fields(msg); len(msgs) > 1 {
			fmt.Println("Print and sent", msgs[0], msgs[1])
			sender.Send(msg, 0)
		}
	}
}

package main

import (
	"fmt"
	"strings"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	defer subscriber.Close()

	subscriber.Connect("tcp://localhost:5556")

	// Note that when you use a SUB socket you must set a subscription using zmq_setsockopt() and SUBSCRIBE.
	// ZMQ_SUBSCRIBE: Establish message filter.

	// See: http://api.zeromq.org/4-1:zmq-setsockopt#toc41.
	// The subscriber can set many subscriptions, which are added together.
	subscriber.SetSubscribe("1")
	subscriber.SetSubscribe("3")

	// The subscriber can also cancel specific subscriptions
	// https://godoc.org/github.com/pebbe/zmq4#Socket.SetUnsubscribe

	for {
		// Just 1 and 3.
		msg, _ := subscriber.Recv(0)

		if msgs := strings.Fields(msg); len(msgs) > 1 {
			fmt.Println("msg", msgs[0], msgs[1], msgs[2])
		}
	}
}

package main

import (
	zmq "github.com/pebbe/zmq4"
)

func main() {
	front, _ := zmq.NewSocket(zmq.ROUTER)
	defer front.Close()
	front.Bind("tcp://*:5550")

	back, _ := zmq.NewSocket(zmq.DEALER)
	defer back.Close()
	back.Bind("tcp://*:5560")

	poller := zmq.NewPoller()
	poller.Add(front, zmq.POLLIN)
	poller.Add(back, zmq.POLLIN)

	for {
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {
			switch s := socket.Socket; s {
			case front:
				for {
					msg, _ := front.Recv(0)
					// More message data parts to follow
					// The ZMQ_RCVMORE option shall return True (1) if the message part last received from the socket was a data part with more parts to follow. If there are no data parts to follow, this option shall return False (0).
					if ok, _ := s.GetRcvmore(); ok {
						back.Send(msg, zmq.SNDMORE)
					} else {
						back.Send(msg, 0)
						break
					}
				}
			case back:
				for {
					msg, _ := back.Recv(0)
					if ok, _ := s.GetRcvmore(); ok {
						front.Send(msg, zmq.SNDMORE)
					} else {
						front.Send(msg, 0)
						break
					}
				}
			}
		}
	}
}

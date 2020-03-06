package main

// Vent

import (
	"strconv"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	sender, _ := zmq.NewSocket(zmq.PUSH)
	defer sender.Close()
	sender.Bind("tcp://*:5557")

	sink, _ := zmq.NewSocket(zmq.PUSH)
	defer sink.Close()
	sink.Connect("tcp://localhost:5558")

	// Give 0MQ time to synchronize the start of the batch with all workers being up and running.
	//
	// This is a fairly common gotcha in ZeroMQ and there is no easy solution. The zmq_connect method takes a certain time. So when a set of workers connect to the ventilator, the first one to successfully connect will get a whole load of messages in that short time while the others are also connecting.
	time.Sleep(2 * time.Second)

	sink.Send("0", 0)

	for i := 0; i < 100; i++ {
		sender.Send(strconv.Itoa(i), 0)
	}
	//  Give 0MQ time to deliver.
	time.Sleep(2 * time.Second)
}

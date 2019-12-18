package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// WaitGroup is used to wait for the program to finish goroutines
var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// wg.Add(2)
	// // With a goroutine, return immediately to the next line and don’t wait for the function to complete
	// go foo()
	// go bar()
	// wg.Wait()

	// Channels provide a way for two goroutines to communicate with each other and syn‐ chronize their execution
	// {string, int...} the things that are passed on the channel
	var c chan string = make(chan string)
	// Using a channel like this synchronizes the two goroutines. When pinger attempts to send a message on the channel, it will wait until printer is ready to receive the message
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

func foo() {
	for i := 0; i < 12; i++ {
		fmt.Println("Foo:", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(amt * time.Millisecond)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 12; i++ {
		fmt.Println("Bar:", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(amt * time.Millisecond)
	}
	wg.Done()
}

// Now pinger is only allowed to send to c
// func pinger(c chan<- string) { ...
func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

// bidirectional channel
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

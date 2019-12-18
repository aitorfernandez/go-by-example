package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	// With a goroutine, return immediately to the next line and don’t wait for the function to complete
	go foo()
	go bar()
	wg.Wait()
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

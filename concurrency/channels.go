package main

import (
	"fmt"
	"os"
	"time"
)

/*
. 	This is a situation when two goroutines wait for each other and non of them can proceed its execution
	Golang can detect deadlocks in runtime that’s why we can see this error.
	This error occurs because of the blocking nature of communication operations.
*/
func deadlock() {
	c := make(chan int)
	c <- 42    // write to a channel
	val := <-c // read from a channel
	fmt.Println(val)
}

// A channel that doesnt specify a direction
// is known as a bi-directional channel
func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

// Channel direction specified
// its restricted to only receive data
func ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

// We can also say that the channel is only defined to be sent
// Attempt to receive from the channel will result in a compiler error
func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		// fmt.Println(<-c) // its ok also
		time.Sleep(time.Second * 1)
	}
}

func startPingPong() {

	var c chan string = make(chan string)

	// start the go routines
	go pinger(c)
	go ponger(c)
	go printer(c)

	// the program will hang here ... waiting for input
	var input string
	fmt.Scanln(&input)
}

func usingSelect() {

	// create two unbuffered channels
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from go routine 1"
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for {
			c2 <- "from go routine 2"
			time.Sleep(time.Second * 5)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Second * 10):
				fmt.Println("timeout")
				os.Exit(1)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}

func main() {
	//startPingPong()
	usingSelect()
	//deadlock()
}

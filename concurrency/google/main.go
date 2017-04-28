package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// oneCall()
	// twoCalls()
	// twoCallsWithFanIn()
	// twoCallsWithSequencing()
	twoCallsWithFanInAndSelect()

	fmt.Println("You're boring; I'm leaving.")
}

func oneCall() {
	c := boring("boring!") // Function returning a channel
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
}

func twoCalls() {
	joe := boring("Joe")
	ann := boring("Ann")
	for i := 0; i < 5; i++ {
		fmt.Printf(<-joe)
		fmt.Printf(<-ann)
	}
}

func twoCallsWithFanIn() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Printf(<-c)
	}
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func twoCallsWithFanInAndSelect() {
	c := fanInWithSelect(boring("Joe"), boring("Ann"))

	timeout := time.After(2 * time.Second)
	for i := 0; i < 10; i++ {
		select {
		case s := <-c:
			fmt.Printf(s)

		// timeout for each message.
		// case <-time.After(1 * time.Second):

		// timeout for entire conversation.
		case <-timeout:
			fmt.Println("You're too slow.")
			return
		}
	}
}

func fanInWithSelect(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()

	return c
}

func boring(msg string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d\n", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c // return the channel to the caller.
}

/////////////////////////////////////////////////////////////////////////
func twoCallsWithSequencing() {

	c := make(chan Message)
	go func() {
		for {
			c <- <-boringWait("Joe")
		}
	}()
	go func() {
		for {
			c <- <-boringWait("Ann")
		}
	}()

	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)

		msg1.wait <- true
		msg2.wait <- true
	}
}

type Message struct {
	str  string
	wait chan bool
}

func boringWait(msg string) <-chan Message { // Returns receive-only channel of strings.

	waitForIt := make(chan bool)
	c := make(chan Message)
	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			<-waitForIt
		}
	}()

	return c // return the channel to the caller.
}

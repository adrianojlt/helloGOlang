package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	// Start a goroutine that reads a value from a channel and prints it
	go func(ch chan int) {
		fmt.Println("start")
		fmt.Println(<-ch)
	}(ch)

	// Start a goroutine that prints a dash every second
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			fmt.Println("-")
		}
	}()

	// Sleep for two seconds
	time.Sleep(2500 * time.Millisecond)

	// Send a value to the channel
	ch <- 5

	// Sleep three more seconds to let all goroutines finish
	time.Sleep(3 * time.Second)
}

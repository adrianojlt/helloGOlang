package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	sum := func(a int, b int) <-chan int {
		ch := make(chan int)
		go func() {
			// Random time up to one second
			delay := time.Duration(r.Int()%1000) * time.Millisecond
			time.Sleep(delay)
			ch <- a + b
			close(ch)
		}()
		return ch
	}

	// Call sum 4 times with the same parameters
	ch1 := sum(3, 5)
	ch2 := sum(3, 5)
	ch3 := sum(3, 5)
	ch4 := sum(3, 5)

	// wait for the first goroutine to write to its channel
	select {
	case result := <-ch1:
		fmt.Printf("ch1: 3 + 5 = %d", result)
	case result := <-ch2:
		fmt.Printf("ch2: 3 + 5 = %d", result)
	case result := <-ch3:
		fmt.Printf("ch3: 3 + 5 = %d", result)
	case result := <-ch4:
		fmt.Printf("ch4: 3 + 5 = %d", result)
	}
}

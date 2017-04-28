package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var counter int
var c = make(chan int)
var done = make(chan bool)
var x int64

func main() {
	start := time.Now()
	go incrementor("Foo:")
	go incrementor("Bar:")
	go puller()
	<-done

	fmt.Println("Final counter: ", counter)

	fmt.Println("Done in ", time.Since(start))
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		c <- 1
		fmt.Println(s, i)
		if i == 19 {
			atomic.AddInt64(&x, 1)
			fmt.Println("XXXXXXXXXXXXX", x)
		}
		if atomic.LoadInt64(&x) == 2 {
			close(c)
		}
	}
}

func puller() {
	for {
		i, more := <-c
		if more {
			counter += i
			fmt.Println("Counter", counter)
		} else {
			done <- true
			close(done)
			return
		}
	}
}

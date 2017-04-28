package main

import (
	"fmt"
	"time"
)

func main() {
	// oneChannel()
	// multipleChannels()
	oneChannelV2()

	fmt.Println("Loop done!")
}

func oneChannel() {
	c := make(chan int)

	go getNumbers(c)

	for number := range c {
		fmt.Println(number)
	}
}

func multipleChannels() {
	even := make(chan int)
	odd := make(chan int)

	go getEvenNumbers(even)
	go getOddNumbers(odd)

	start := time.Now()

	for {
		select {
		case num := <-even:
			fmt.Printf("even: %2d %5d\n", num, msSince(start))
		case num := <-odd:
			fmt.Printf("odd:  %2d %5d\n", num, msSince(start))
		}

	}
}

func oneChannelV2() {
	c, quit := getNumbersV2()

	timeout := time.After(3 * time.Second)

loop:
	for {
		select {
		case number := <-c:
			fmt.Println(number)
		// if number > 10 {
		case <-timeout:
			quit <- true
			break loop
		}
	}
}

//////////////////////////////////////////////////////////////////////////////

func getNumbers(c chan int) {
	for i := 1; i <= 10; i++ {
		c <- i
		time.Sleep(250 * time.Millisecond)
	}

	close(c)
}

func getEvenNumbers(c chan int) {
	for i := 2; ; i += 2 {
		c <- i
		time.Sleep(500 * time.Millisecond)
	}
}

func getOddNumbers(c chan int) {
	for i := 1; ; i += 2 {
		c <- i
		time.Sleep(1 * time.Second)
	}
}

func msSince(start time.Time) int64 {
	return int64(time.Since(start) / time.Millisecond)
}

func getNumbersV2() (c chan int, quit chan bool) {
	c = make(chan int)
	quit = make(chan bool)

	go func() {
		for i := 1; ; i++ {
			select {
			case c <- i:
				time.Sleep(125 * time.Millisecond)
			case <-quit:
				return
			}
		}
	}()

	return c, quit
}

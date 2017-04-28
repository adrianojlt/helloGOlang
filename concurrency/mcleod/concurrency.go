package main

import "fmt"
import "sync"
import "time"

var wg sync.WaitGroup

// Parallelism
// func init() {
// 	runtime.GOMAXPROCS(runtime.NumCPU())
// }

func main() {
	start := time.Now()

	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()

	fmt.Println("Done in ", time.Since(start))
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo: ", i)
		time.Sleep(time.Duration(10 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar: ", i)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
	wg.Done()
}

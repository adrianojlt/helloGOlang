package main

import (
	"fmt"
)

func log(msg string) {

}

func add(a int, b int) int {
	return a + b
}

func power(name string) (int, bool) {
	return 0,true;
}

func anonymousFunctions() {

	// anonymous functions with different type returns
	a,b := func(a,b int,s string) (int,int) {
		return 2,3
	}(1,2,"string")

	a,b = func(a bool,b int) (int,int) {
		return 2,3
	}(true,120)

	fmt.Printf("%d , %d\n",a,b)
}


func main() {
	_, exists := power("cenadas")

	if ( exists ) {}

	func() {
		fmt.Println("anonymous function")
	}()

}
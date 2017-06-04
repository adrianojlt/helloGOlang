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
	return 0, true
}

func anonymousFunctions() {

	// anonymous functions with different type returns
	a, b := func(a, b int, s string) (int, int) {
		return 2, 3
	}(1, 2, "string")

	a, b = func(a bool, b int) (int, int) {
		return 2, 3
	}(true, 120)

	fmt.Printf("%d , %d\n", a, b)
}

// named return values
func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value
}

// a function like this together with a nonlocal variable it
// references is known as a closure
func closureExample01() {
	x := 0

	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
}
func closureExample02() {
	makeEvenGenerator := func() uint {

	}
}

func main() {
	//test()
	//closureExample01()
	closureExample02()
}

func test() {
	_, exists := power("cenadas")

	if exists {
	}

	func() {
		fmt.Println("anonymous function")
	}()
}

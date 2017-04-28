package main

import (
	"fmt"
)

func init() {
}

func main() {
	fmt.Println("initializers")
	initializers()
	fmt.Println("\nslicing")
	slicing()
}

func initializers() {

	initSlice := []string{"a","b","c"}
	fmt.Println(initSlice)

	a := make([]int,5,10)
	b := []int{0,0,0,0,0}

	fmt.Println("a = ", a)
	fmt.Println("len(a) = ", len(a))
	fmt.Println("cap(a) = ", cap(a))

	fmt.Println("b = ", b)
	fmt.Println("len(b) = ", len(b))
	fmt.Println("cap(b) = ", cap(b))
}

func slicing() {

	g := []byte{'g', 'o', 'l', 'a', 'n', 'g'}

	fmt.Println(string(g[:])) 		// golang
	fmt.Println(string(g[1:4])) 	// ola
	fmt.Println(string(g[2:])) 		// lang
}


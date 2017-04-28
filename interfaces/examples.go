package main

import "fmt"
import "os"

type World struct{}

func (w *World) String() string {
	return "new world"
}

func main() {
	fmt.Printf("hello, %s\n", "world")
	fmt.Printf("hello, %s\n", new(World))
	fmt.Fprintf(os.Stdout, "hello to Stand")
}

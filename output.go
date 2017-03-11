package main

import "fmt"

func main() {

	plnBytes, err := fmt.Println("Println");

    fmt.Printf("%d", plnBytes)
    fmt.Fprintf("%s", err)
}
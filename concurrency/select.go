package main

import (
	"fmt"
	"strings"
	"time"
)

func readword(ch chan string) {
	fmt.Println("Type a word, then hit Enter.")
	var word string
	fmt.Scanf("%s", &word)
	ch <- word
}

func timeout(t chan bool) {
	time.Sleep(5 * time.Second)
	t <- true
}

func main() {
	selectFunc01()
}

func selectFunc01() {

	t := make(chan bool)
	go timeout(t)

	ch := make(chan string)
	go readword(ch)

	select {
	case word := <-ch:
		fmt.Println("Received", word)
	case <-t:
		fmt.Println("Timeout.")
	}
}

// do the same but with a loop
func selectFunc02() {

	var op string

	for strings.Compare("exit", op) != 0 {
	}
}

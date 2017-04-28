package main

import (
	"fmt"
	"net/http"
	"time"
)

type User struct {
	Cash int
}

func (u *User) sendCash(to *User, amount int) bool {
	if u.Cash < amount {
		return false
	}

	/* Delay to demonstrate the race condition */
	time.Sleep(500 * time.Millisecond)

	u.Cash = u.Cash - amount
	to.Cash = to.Cash + amount
	return true
}

func main() {
	me := User{Cash: 500}
	you := User{Cash: 500}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		me.sendCash(&you, 50)
		fmt.Fprintf(w, "I have $%d\n", me.Cash)
		fmt.Fprintf(w, "You have $%d\n", you.Cash)
		fmt.Fprintf(w, "Total transferred: $%d\n", (you.Cash - 500))
	})

	http.ListenAndServe(":8888", nil)
}

package main

import (
	"fmt"
	"net/http"
	"time"
)

type User struct {
	Cash int
}

type Transfer struct {
	Sender    *User
	Recipient *User
	Amount    int
}

//var transferLock *sync.Mutex

func sendCashHandler(transferchan chan Transfer) {
	var val Transfer
	for {
		val = <-transferchan
		val.Sender.sendCash(val.Recipient, val.Amount)
	}
}

func (u *User) sendCash(to *User, amount int) bool {
	//transferLock.Lock()

	/* Defer runs this function whenever sendCash exits */
	//defer transferLock.Unlock()

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

	transferchan := make(chan Transfer)
	go sendCashHandler(transferchan)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		transfer := Transfer{Sender: &me, Recipient: &you, Amount: 50}
		transferchan <- transfer
		fmt.Fprintf(w, "I have $%d\n", me.Cash)
		fmt.Fprintf(w, "You have $%d\n", you.Cash)
		fmt.Fprintf(w, "Total transferred: $%d\n", (you.Cash - 500))
	})

	http.ListenAndServe(":8888", nil)
}

package main

import "fmt"

type User struct {
	name 		string
	email		string
	ext			int
	enabled		bool
}

func (u *User) setName(name string) {
	u.name = name
}

func init() {
	fmt.Println("init() is called before main");
}

func main() {
	arrays();
}

func arrays() {

	var john User
	bill := new(User)
	bill.name = "BILL"
	fmt.Println(john.enabled)
	fmt.Println(bill.name)

	var initArray = []string {"2","3"}
	fmt.Println(initArray);
}

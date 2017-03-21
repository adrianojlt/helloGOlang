package main

import "fmt"

type User struct {
	name    string
	email   string
	ext     int
	enabled bool
}

type notExported struct { //this struct is visible only in this package as it starts with small letter

}

//variable starts with capital letter, so visible outside this package
type Exported struct {
	notExportedVariable int    // variable starts with small letter, so NOT visible outside package
	ExportedVariable    int    // variable starts with capital letter, so visible outside package
	s                   string // not exported
	S                   string // exported
}

type Rectangle struct {
	length, width int
	name          string
}

func (u *User) setName(name string) {
	u.name = name
}

func initRectangle() {
	r1 := Rectangle{2, 1, "my_r1"} // initialize values in order they are defined in struct
	fmt.Println("Rectangle r1 is: ", r1)

	r2 := Rectangle{width: 3, name: "my_r2", length: 4} // initialize values by variable name in any order
	fmt.Println("Rectangle r2 is: ", r2)

	pr := new(Rectangle) // get pointer to an instance with new keyword
	(*pr).width = 6      // set value using . notation by dereferencing pointer.
	pr.length = 8        // set value using . notation - same as previous.  There is no -> operator like in c++. Go automatically converts

	pr.name = "ptr_to_rectangle"
	fmt.Println("Rectangle pr as address is: ", &pr)
	fmt.Println("Rectangle pr as value is: ", pr)  // Go performs default printing of structs
	fmt.Println("Rectangle pr as value is: ", *pr) // address and value are differentiated with an & symbol
}

func init() {
	fmt.Println("init() is called before main")
}

func main() {
	arrays()
	initRectangle()
}

func arrays() {

	var john User
	bill := new(User)
	bill.name = "BILL"
	fmt.Println(john.enabled)
	fmt.Println(bill.name)

	var initArray = []string{"2", "3"}
	fmt.Println(initArray)
}

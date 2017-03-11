package main

// fmt package provided by the standard library.
// format and output data
import "fmt"


func prints() {
	hello := "hello"
	fmt.Println(hello + " Println")
	fmt.Printf(hello + " Printf")
	fmt.Sprintf(hello + " Sprintf")
	fmt.Fprint(hello + " Fprint")
	fmt.Fprintf(hello + " Fprintf")
	fmt.Fprintln(hello + " Fprintln")
	fmt.Print(hello + " Print")
	fmt.Sprint(hello + " Sprint")
	fmt.Sprintln(hello + " Sprintln")
	fmt.Errorf(hello + " Errorf")
}

func main() {

	prints()

	count := 100;
	count = 0;
	name, power := "Adriano", 9000
    fmt.Printf("hello, world\n")
    fmt.Printf("name %s and power %d and count = %d \n", name, power, count)
    fmt.Printf("Printf\n")
    fmt.Println("hello, world\n")
    fmt.Sprintf("Sprintf\n")
}


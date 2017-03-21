package main  

import (
	"fmt"
)

type Shaper interface {
	Area() int
}

type Rectangle struct {
	length, width int
}

type Square struct {
	side int
}

// this function has the same signature as the Shaper interface
// Rectangle now implements the Shaper interface
func (r Rectangle) Area() int {
	return r.length * r.width
}

// Like in other languages (like Java and C#) the declaration of implementation
// is explicit. In GO this is not necessary.
// If a type has a function with the same interface signature means that type
// implements that interface
func (s Square) Area() int {
	return s.side * s.side
}

func iterate() {

	shapers := [...]Shaper{
		Rectangle{length:5,width:2},
		Rectangle{length:10,width:22},
		Square{side:53},
		Square{side:5}}

	fmt.Println("\nLooping through shapes for area ...")

   	for n, _ := range shapers {
       fmt.Println("Shape details: ", shapers[n])
       fmt.Println("Area of this shape is: ", shapers[n].Area())
   	}
}

func main() {

	var s Shaper

	re := Rectangle{length:5,width:3}
	sq := Square{side:5}

	fmt.Println("Rectangle r details are: ", re)  
	fmt.Println("Square sq details are: ", sq)  

  	fmt.Println("Rectangle r's area is: ", re.Area())  
  	fmt.Println("Square r's area is: ", sq.Area())  

  	s = Shaper(re)
   	fmt.Println("Area of the Shape Rectangle is: ", s.Area())  
   	s = sq
   	fmt.Println("Area of the Shape Squere is: ", s.Area())  

   	iterate()
}


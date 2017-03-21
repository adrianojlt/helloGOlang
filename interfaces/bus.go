package main

import "fmt"

type Bus struct {
    l, b, h int
    rows, seatsPerRow int
}

type Cuboider interface {
    CubicVolume() int
}

// 
func (b Bus) CubicVolume() int {
    return b.l * b.b * b.h
}

type PublicTransporter interface {
    PassengerCapacity() int
}

// Go Interfaces are data centric. You define the data first and the interfaces 
// abstractions are built as you go along
func (b Bus) PassengerCapacity() int {
    return b.rows * b.seatsPerRow
}

func main() {
    b := Bus{l:10, b:6, h:3, rows:10, seatsPerRow:5}

    fmt.Println("Cubic volume of b:", b.CubicVolume())
    fmt.Println("Maximum number of passengers:", b.PassengerCapacity())
    fmt.Println("Is compliant with law:", b.IsCompliantWithLaw())
}

// Go's interfaces work to find out how it could aid better extensibility and evolution ...
type PersonalSpaceLaw interface {
    IsCompliantWithLaw() bool
}

// ... this way you dont need to mess arround 'Bus' type, you just need to 
// create a function with the same signature as the method interface
func (b Bus) IsCompliantWithLaw() bool {
    return (b.l * b.b * b.h) / (b.rows * b.seatsPerRow) >= 3
}

package main

import (
	"fmt"
)

func init() {
}

func main() {
	initializers()
	arrays()
}

func initializers() {
	var notUsedArr [5]int
	_ = notUsedArr // avoid the 'declared and not used' warning
	
	iniArrWithSize := [2]string{"first","secound"}
	iniArrWithoutSize := [...]string{"fisrt","secound"}

	Use(iniArrWithSize,iniArrWithoutSize)
}

func arrays() {
	var initArray = []string{"2", "3"}
	fmt.Println(initArray)
}

func Use(vals ...interface{}) {
    for _, val := range vals {
        _ = val
    }
}

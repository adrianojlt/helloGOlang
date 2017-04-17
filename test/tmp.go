package main

import (
	//"bufio"
    //"fmt"
    "io/ioutil"
    //"os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

	const BOM  = string('\uFEFF')

	//i := BOM + " cenas é para quê"
	//d1 := []byte("cenas é para quê")
	str := "\uFEFF quê"
	//d1 := []byte("\uFEFF cenas é para quê")
	err := ioutil.WriteFile("tmp.txt", []byte(str), 0644)
    check(err)
	//fmt.Println(BOM)

    /*
	f, err := os.Create("/tmp/dat2")
    check(err)

    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)

     n3, err := f.WriteString("writes\n")
    fmt.Printf("wrote %d bytes\n", n3)

     f.Sync()

    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    fmt.Printf("wrote %d bytes\n", n4)

    w.Flush()
    */
	
}
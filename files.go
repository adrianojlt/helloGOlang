package main

import "fmt"
import "log"
import iou "io/ioutil"

func main() {

	fmt.Println("List of files: ")

	files, err := iou.ReadDir("testdir")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

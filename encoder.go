package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"golang.org/x/text/encoding/charmap"
)

func create() {
	iFile, err := os.Open("isotext.txt")
	if err != nil {
		// ... file open error
	}

	oFile, err := os.Create("utf8.txt")
	if err != nil {
		// ... file open error
	}

	r := charmap.ISO8859_1.NewDecoder().Reader(iFile)
	io.Copy(oFile,r)

	oFile.Close();
	iFile.Close()
}

func BOMit(filePath string) {
	// Open a file.
    f, _ := os.Open(filePath)

    // Use bufio.NewReader to get a Reader.
    // ... Then use ioutil.ReadAll to read the entire content.
    reader := bufio.NewReader(f)
    content, _ := ioutil.ReadAll(reader)

    // The UTF-8 BOM is a sequence of bytes (EF BB BF) that allows the reader to identify a file as being encoded in UTF-8.
    // The byte order mark (BOM) is a Unicode character, U+FEFF
    // at the start of a text stream can signal several things to a program consuming the text
    total := "\uFEFF" + string(content)
    f.Close()

    // Print File content.
    fmt.Println(string(content))

    // write ...
    ioutil.WriteFile(filePath, []byte(total), 0644)
}

func main() {
	//create()
	BOMit("test.txt")
}


package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding/charmap"
)

var extensions = []string{"srt", "sub", "txt"}

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
	io.Copy(oFile, r)

	oFile.Close()
	iFile.Close()
}

func readFile(path string) []byte {
	// Open a file.
	f, _ := os.Open(path)

	// Use bufio.NewReader to get a Reader.
	// ... Then use ioutil.ReadAll to read the entire content.
	reader := bufio.NewReader(f)
	originalContent, _ := ioutil.ReadAll(reader)

	f.Close()
	return originalContent
}

func convert(originalContent []byte) []byte {
	var convertedContent []byte

	// try to detect the encoding used
	detector := chardet.NewTextDetector()
	result, _ := detector.DetectBest(originalContent)
	charset := result.Charset

	if strings.Contains(charset, "UTF") {
		convertedContent = originalContent
	} else {
		// now ... here i am assuming ISO08859 but i need to use the one found in the previous detect encoding operation
		convertedContent, _ = charmap.ISO8859_1.NewDecoder().Bytes(originalContent)
	}

	return convertedContent
}

func boomIT(content []byte) {

	// The UTF-8 BOM (Byte Order Mark) is a sequence of bytes (EF BB BF) that allows the reader to identify a file as being encoded in UTF-8.
	// The byte order mark (BOM) is a Unicode character, U+FEFF
	// at the start of a text stream can signal several things to a program consuming the text
	total := "\uFEFF" + string(convertedContent)

	// write (default encoding in golang is UTF-8) ...
	ioutil.WriteFile(filePath, []byte(total), 0644)
}

func bomIt(filePath string, showContent bool) {

	var convertedContent []byte

	// Open a file.
	f, _ := os.Open(filePath)

	// Use bufio.NewReader to get a Reader.
	// ... Then use ioutil.ReadAll to read the entire content.
	reader := bufio.NewReader(f)
	originalContent, _ := ioutil.ReadAll(reader)

	// try to detect the encoding used
	detector := chardet.NewTextDetector()
	result, _ := detector.DetectBest(originalContent)
	charset := result.Charset

	if strings.Contains(charset, "UTF") {
		convertedContent = originalContent
	} else {
		// now ... here i am assuming ISO08859 but i need to use the one found in the previous detect encoding operation
		convertedContent, _ = charmap.ISO8859_1.NewDecoder().Bytes(originalContent)
	}

	// The UTF-8 BOM (Byte Order Mark) is a sequence of bytes (EF BB BF) that allows the reader to identify a file as being encoded in UTF-8.
	// The byte order mark (BOM) is a Unicode character, U+FEFF
	// at the start of a text stream can signal several things to a program consuming the text
	total := "\uFEFF" + string(convertedContent)
	f.Close()

	// write (default encoding in golang is UTF-8) ...
	ioutil.WriteFile(filePath, []byte(total), 0644)
}

func walker(path string, f os.FileInfo, err error) error {

	if !f.IsDir() && strings.HasSuffix(f.Name(), "srt") {
		boomIT(convert(readFile(path)))
		//bomIt(path, true)
	}

	return nil
}

func list(path string, f os.FileInfo, err error) error {

	if !f.IsDir() && strings.HasSuffix(f.Name(), "srt") {
		fmt.Println(f.Name(), "\t encoded with: ", detect(path))
	}

	return nil
}

func detect(path string) string {
	f, _ := os.Open(path)
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)
	detector := chardet.NewTextDetector()
	result, _ := detector.DetectBest([]byte(content))
	return result.Charset
}

func start() {
	flag.Parse()
	path := flag.Arg(0)
	err := filepath.Walk(path, walker)
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func main() {
	//create()
	start()
	//test()
	//bomIt("./testdir/iso8859.srt", false)
}

func test() {
	flag.Parse()
	folder := flag.Arg(0)
	_ = filepath.Walk(folder, list)
	//name := charmap.ISO8859_1.String()
}

/*
try this
boomit --dir="./dir"
boomit --dir="c:\temp\subtitles" -x
*/
func args() {

}

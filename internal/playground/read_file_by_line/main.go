package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Helper function to check for errors
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Read file contents into memory
	dat, err := ioutil.ReadFile("text.txt")
	check(err)
	fmt.Print(string(dat) + "\n")

	// Open() takes the name of the file to open as an argument and returns a
	// file handle (or object) and error if there is any. When we open a file
	// using Open() function call, it opens in a read-only mode. To be able to
	// write to a file, we would have to use the OpenFile() function.
	f, err := os.Open("text.txt")
	check(err)

	// Read first 5 bytes
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// Use Seek to read from a certain place
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	// Or use ReadAtLeast
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 8)
	n3, err := io.ReadAtLeast(f, b3, 8)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// "Rewind" using Seek
	_, err = f.Seek(0, 0)
	check(err)

	// Use a buffered reader for its efficiency with small reads and additional
	// options
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	for idx, eachline := range txtlines {
		fmt.Printf("%d | %s\n", idx, eachline)
	}

	// Close the file when done
	err = f.Close()
	check(err)
}

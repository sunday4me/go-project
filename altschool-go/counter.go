package main

import (
	"fmt"
	"io"
	"os"
)

//var i io.Reader
//var w io.Writer

type Counter int

func (c *Counter) Write(b []byte) (int, error) {
	l := len(b)
	*c += Counter(l) //casting l to an integer so it becomes the right type
	return l, nil
}

func main() {
	var c Counter
	f1, _ := os.Open("input.txt")
	//f2, _ := os.Create("output.txt")
	newWriter := &c
	n, _ := io.Copy(newWriter, f1)

	fmt.Println("Copied", n, "bytes")
	//fmt.Println(c)

}

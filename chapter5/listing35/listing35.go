// Sample program to show how a bytes.Buffer can also be used
// with the io.Copy function.
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// main is the entry point for the application.
func main() {
	var b bytes.Buffer

	//slice :=[]byte("Hello")
	//println(slice[0])

	// Write a string to the buffer.
	b.Write([]byte("Hello")) //可以这么定义切片么？

	// Use Fprintf to concatenate a string to the Buffer.
	fmt.Fprintf(&b, "World!")

	// Write the content of the Buffer to stdout.
	io.Copy(os.Stdout, &b)
}

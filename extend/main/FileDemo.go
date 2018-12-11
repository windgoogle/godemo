package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	buf := "Hello, World"
	//os.Setenv("TMP","E:\\test")
	os.Chdir("F:")
	file, err := ioutil.TempFile(".", "tmpfile")
	if err != nil {
		panic(err)
	}
	defer os.Remove(file.Name())

	if _, err := file.Write([]byte(buf)); err != nil {
		panic(err)
	}

	fmt.Println(file.Name())

}

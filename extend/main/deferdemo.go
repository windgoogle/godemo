package main

import (
	"fmt"
	"os"
)

/**
多个defer ,最先声明的，最后执行
*/
func main() {
	fsrc, err := os.Open("E:\\goproject\\src\\goinaction\\extend\\main\\source.txt")
	if err != nil {
		fmt.Println("open source file failed")
		return
	}

	//defer fsrc.Close()
	defer func() {
		fmt.Println("close source file!")
		fsrc.Close()
	}()

	fdes, err := os.Open("E:\\goproject\\src\\goinaction\\extend\\main\\target.txt")
	if err != nil {
		fmt.Println("open target file failed")
		return
	}

	//defer fdes.Close()
	defer func() {
		fmt.Println("close target file !")
		fdes.Close()
	}()
	fmt.Println("do smomething here")

}

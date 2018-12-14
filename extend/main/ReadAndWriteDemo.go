package main

import (
	"fmt"
	"io"
)

//没觉得这个devNull类型很奇怪么
//类型基于基本类型
//却实现了io.writer接口的方法，可以赋值io.writer变量
type devNull int

func (devNull) Write(p []byte) (int, error) {
	return len(p), nil
}

func main() {
	c := devNull(10)
	fmt.Println(c)
	var Dis io.Writer = devNull(0)
	fmt.Println(Dis)
}

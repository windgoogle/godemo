package main

import (
	"errors"
	"fmt"
	"os"
)

/*
  GO语言中错误是个接口
   type error interface {
      Error() string
    }
   这个小写的error却是全局可见的

   Go 语言内置了一个通用错误类型，在 errors 包里。这个包还提供了一个 New() 函数让我们方便地创建一个通用错误
	package errors

	func New(text string) error {
    	return &errorString{text}
	}

	type errorString struct {
    	s string
	}

	func (e *errorString) Error() string {
    	return e.s
	}

    如果你的错误字符串需要定制一些参数，可使用 fmt 包提供了 Errorf 函数
	var thing = "something"
	var err = fmt.Errorf("%s happened", thing)
*/

type SomeError struct {
	Reason string
}

func (s SomeError) Error() string {
	return s.Reason
}

func main() {
	var err1 error = SomeError{"something happened"}
	fmt.Println(err1)
	var err2 = errors.New("something happened") //由 errors包里的New函数产生
	fmt.Println(err2)

	var thing = "something"
	var err3 = fmt.Errorf("%s happened", thing)
	fmt.Println(err3)

	// 打开文件
	var f, err = os.Open("main.go")
	if err != nil {
		// 文件不存在、权限等原因
		fmt.Println("open file failed reason:" + err.Error())
		return
	}
	// 推迟到函数尾部调用，确保文件会关闭
	defer f.Close()
	// 存储文件内容
	var content = []byte{}
	// 临时的缓冲，按块读取，一次最多读取 100 字节
	var buf = make([]byte, 100)
	for {
		// 读文件，将读到的内容填充到缓冲
		n, err := f.Read(buf)
		if n > 0 {
			// 将读到的内容聚合起来
			content = append(content, buf[:n]...)
		}
		if err != nil {
			// 遇到流结束或者其它错误
			break
		}
	}
	// 输出文件内容
	fmt.Println(string(content))

}

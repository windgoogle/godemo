package main

import "fmt"

var negErr = fmt.Errorf("non positive number")

func main() {
	defer func() {
		if err := recover(); err != nil { //recover catch 住了函数
			//fmt.Println("error catched", err)
			if err == negErr {
				fmt.Println("error catched", err) //catch
			} else {
				panic(err) //rethrow
			}
		}
	}() //括号表示匿名函数的调用
	fmt.Println(fact(10))
	fmt.Println(fact(5))
	fmt.Println(fact(-5))
	fmt.Println(fact(15))
}

// 让阶乘函数返回错误太不雅观了
// 使用 panic 会合适一些
func fact(a int) int {
	if a <= 0 {
		panic(negErr)
	}
	var r = 1
	for i := 1; i <= a; i++ {
		r *= i
	}
	return r
}

package main

import "fmt"

func ExFunc(n int) func() {
	sum := n
	return func() { //把匿名函数作为值赋给变量a (Go 不允许函数嵌套。
		//然而你可以利用匿名函数实现函数嵌套)
		fmt.Println(sum + 1) //调用本函数外的变量
	} //这里没有()匿名函数不会马上执行
}

func main() {
	myFunc := ExFunc(10)
	myFunc() // 11
	myAnotherFunc := ExFunc(20)
	myAnotherFunc() //21
	myFunc()        //11
	myAnotherFunc() //21
	fmt.Printf("The quick brown fox jumped over lazy dogs", 3.14)

}

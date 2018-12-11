package main

import "fmt"

/**
组合继承接口的例子
*/
type Smellable1 interface {
	smell()
}

type Eatable1 interface {
	eat()
}

type Fruitable1 interface {
	Smellable1
	Eatable1 //没有声明类型，匿名的，外层接口拥有其所有方法 （接口和结构体均可以）
}

//接口变量赋值的例子
//变量赋值的本质就是一次内存浅拷贝

type Rect struct {
	Width  int
	Height int
}

func main() {
	var a interface{}
	var b interface{}
	var r = Rect{50, 50}
	a = r
	b = &r             //赋值指针
	var rx = a.(Rect)  //类型转换
	var bx = b.(*Rect) //转换成指针类型
	r.Width = 100
	r.Height = 100
	fmt.Println(rx) //浅拷贝，r 的未改变a的值，所以证明是浅拷贝

	fmt.Println(bx) //无拷贝，，内存与指针变量指向的内存共享
}

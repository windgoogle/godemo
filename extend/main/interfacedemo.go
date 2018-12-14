package main

import (
	"fmt"
	"unsafe"
)

type Smellable interface {
	smell()
}

type Eatable interface {
	eat()
}

type Apple struct {
}

//有接受者的函数叫方法
func (a Apple) smell() {
	fmt.Println("apple can smell")

}

func (a Apple) eat() {
	fmt.Println("apple can eat")
}

type Flower struct {
}

func (f Flower) smell() {
	fmt.Println("flower can smell")
}

func main() {
	var s1 Smellable
	var s2 Eatable
	var apple = Apple{}
	var flower = Flower{}
	s1 = apple

	s1.smell()
	s1 = flower
	s1.smell()

	s2 = apple
	s2.eat()
	//空接口 interface{} 任何结构体隐式的实现了空接口
	//GO自己内置了空接口，避免用户重复实现空接口
	//空接口可认为类似java 里的Object类型，任何对象都继承自Object ,就是Any类型
	//这个map 里的值为空接口类型，可以容纳任何类型的对象，在获取需要做类型转换
	var user = map[string]interface{}{
		"age":     30,
		"address": "Beijing Chaoyang",
		"married": true,
	}
	fmt.Println(user)
	//类型转换
	var age = user["age"].(int) //接口.(转换的类型)
	var address = user["address"].(string)
	var married = user["married"].(bool)
	fmt.Println(age, address, married)

	//接口的本质
	//接口可以看做是个特殊的容器，这个容器只能容纳一个对象，只有实现了这个接口类型的对象才可以放进去
	//GO源码中，接口变量也是由结构体来定义的，包含两个指针字段，一个指针对象指向被容纳的对象内存，另一个字段指向一个特殊的结构体itab
	//这个特殊的结构体包含了接口的类型信息和被容纳对象的数据类型信息
	/*
		  type iface strunt{
		      tab #itab       //类型指针
		      data unsafe.Pointer   //数据指针
		   }

		type itab struct {
		inter *interfacetype    //接口类型信息

		_type *_type      //数据类型信息
		 ......
	*/

	var s interface{}
	fmt.Println(unsafe.Sizeof(s)) //结果是16=2个机器字
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(unsafe.Sizeof(arr)) //结果是80=10个机器字
	s = arr
	fmt.Println(unsafe.Sizeof(s)) //给接口变量赋值后，结果还是16=2个机器字，丝毫未影响接口变量的内存占用

}

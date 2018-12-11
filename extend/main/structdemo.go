package main

import (
	"fmt"
	"unsafe"
)

type Point struct {
	x int
	y int
}

type Circle struct {
	x      int
	y      int
	Radius int
}

func (p Point) show() {
	fmt.Println(p.x, p.y)
}

type CircleB struct {
	Radius int
	Point  //匿名结构体
}

func main() {
	var c1 Circle = Circle{
		x:      100,
		y:      100,
		Radius: 50, // 注意这里的逗号不能少
	}
	fmt.Printf("%+v\n", c1)

	var c2 Circle = Circle{
		x: 20,
	}

	fmt.Printf("%+v\n", c2)

	var c3 Circle = Circle{20, 30, 40} //这种创建形式必须给所有初始值，这种创建形式叫做"顺序形式"
	fmt.Printf("%+v\n", c3)

	var c *Circle = &Circle{100, 100, 50} //结构体指针，这里创建了一个结构体，创建了一个指针指向这个结构体
	fmt.Printf("%+v\n", c)

	var c4 *Circle = new(Circle) //new() 函数返回的是指针类型
	fmt.Printf("%+v\n", c4)
	//比较 nil 和零值占用内存的区别，nil 占用一个指针的占用量即一个机器字，而零值不同占用三个机器字
	//另外32位机器和64位机器还有差别，32位只会占用12个字节
	var c5 *Circle = nil
	var c6 Circle

	fmt.Println(unsafe.Sizeof(c5), unsafe.Sizeof(c6))

	fmt.Println("结构体赋值，注意和指针类型的区别")
	//结构体赋值也是浅拷贝，注意和指针类型的区别
	var c7 Circle = Circle{Radius: 50}
	var c8 Circle = c7
	fmt.Printf("%+v\n", c7)
	fmt.Printf("%+v\n", c8)
	c7.Radius = 100
	fmt.Printf("%+v\n", c7)
	fmt.Printf("%+v\n", c8)

	var c9 *Circle = &Circle{Radius: 50}
	var c10 *Circle = c9
	fmt.Printf("%+v\n", c9)
	fmt.Printf("%+v\n", c10)
	c9.Radius = 100
	fmt.Printf("%+v\n", c9)
	fmt.Printf("%+v\n", c10)

	/**
		   GO语言内置的高级数据结构都是有结构体来完成的
		    例如切片
				type slice struct {
	  				array unsafe.Pointer  // 底层数组的地址
	  				len int // 长度
		            cap int // 容量
				}
	*/

	//匿名内嵌结构体
	//内嵌的结构体不提供名称。
	// 这时外面的结构体将直接继承内嵌结构体所有的内部字段和方法，
	// 就好像把子结构体的一切全部都揉进了父结构体一样。
	// 匿名的结构体字段将会自动获得以结构体类型的名字命名的字段名称
	//GO语言通过组合内嵌实现继承
	//Go 语言的结构体明确不支持这种形式的多态，外结构体的方法不能覆盖内部结构体的方法

	var c11 = CircleB{
		Point: Point{
			x: 100,
			y: 100,
		},
		Radius: 50,
	}
	fmt.Printf("%+v\n", c11)
	fmt.Printf("%+v\n", c11.Point)
	fmt.Printf("%d %d\n", c11.x, c11.y) // 继承了字段
	fmt.Printf("%d %d\n", c11.Point.x, c11.Point.y)
	c11.show()       // 继承了方法，如果没有下面注解掉的方法就是继承了
	c11.Point.show() //另一个方面说明，内嵌的结构体方法没有被外层结构体的同方法覆盖掉

}

/*
func (p CircleB) show() {
	fmt.Println("外部结构体show方法")
}
*/

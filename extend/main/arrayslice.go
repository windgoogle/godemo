package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	//a[101]=105    //编译报错， 表示数组长度是编译时就确定的，Invalid array index 101 (out of bounds for 5-element array),
	//如果下标是个变量，又是如何检查的呢，Go会在编译后进行下标检查
	//var b=101
	//a[b]=122    //报错：panic: runtime error: index out of range

	fmt.Println(a)

	//数组赋值，只有子元素类型和数组长度完全相同的数组才能相互赋值
	//赋值的两个数组的元素不是共享的，所以赋值有代价
	var b = a
	fmt.Println("数组赋值后......")
	fmt.Println(a, len(a), cap(a))
	fmt.Println(b, len(b), cap(b))

	a[0] = 1233

	fmt.Println("修改一个数组的值之后.....")
	fmt.Println(a, len(a), cap(a))
	fmt.Println(b, len(b), cap(b))
	fmt.Println("切片的例子.....")
	//切片
	//切边变量， 指针  ， length  ，capacity
	//切片创建

	var s1 []int = make([]int, 5, 8)
	var s2 []int = make([]int, 8) //满容切片

	fmt.Println(s1)
	fmt.Println(s2)

	//切片初始化

	var s3 []int = []int{1, 2, 3, 4, 5}
	fmt.Println(s3, len(s3), cap(s3))

	//空切片，长度和容量均为0 的切片

	var s4 []int = []int{}
	var s5 []int = make([]int, 0)
	var s6 []int //这种又叫nil切片

	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(s6)

	//切片赋值，两个切片赋值，底层数组是共享的，所以对一个切片的改变，会改变另一个，这个和数组不同

	var s7 []int = make([]int, 5, 8)
	for i := 0; i < len(s7); i++ {
		s7[i] = i + 1
	}
	var s8 = s7

	fmt.Println(s7, len(s7), cap(s7))
	fmt.Println(s8, len(s8), cap(s8))

	//尝试修改切片的内容

	s7[0] = 255

	fmt.Println(s7)
	fmt.Println(s8)

	s9 := s7[2:4]
	fmt.Println(s9)
	fmt.Println()
	fmt.Println(s7, len(s7), cap(s7))
	fmt.Println("切割后，子切片容量为什么变为6") //母子依旧共享底层数组，只不过子切片起始指针变化了，
	fmt.Println(s9, len(s9), cap(s9))
	fmt.Println("多种切割方式") //母子依旧共享底层数组，只不过子切片起始指针变化了，
	var t1 = []int{1, 2, 3, 4, 5, 6, 7}
	var t2 = t1[:5]
	var t3 = t1[3:]
	var t4 = t1[:]
	fmt.Println(t1, len(t1), cap(t1))
	fmt.Println(t2, len(t2), cap(t2))
	fmt.Println(t3, len(t3), cap(t3))
	fmt.Println(t4, len(t4), cap(t4))

	//数组变切片
	fmt.Println("数组变切片")
	var c = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //定义数组
	var d = c[2:6]                                 //数组变切片
	fmt.Println(d)
	c[4] = 100 //修改数组后的元素
	fmt.Println(d)

	//copy 函数
	fmt.Println("使用copy函数")
	var s = make([]int, 5, 8)
	for i := 0; i < len(s); i++ {
		s[i] = i + 1
	}

	fmt.Println(s)
	fmt.Println()
	var des = make([]int, 2, 6) //拷贝len 小值
	var n = copy(des, s)
	fmt.Println(n, des)

	//切片的扩容点

}

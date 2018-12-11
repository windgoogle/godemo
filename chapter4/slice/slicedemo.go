package main

import "fmt"

func test(s string) string {
	// 创建一个字符串切片
	// 其长度和容量都是 5 个元素
	slice := make([]string, 5)
	s1 := append(slice, "a")
	println(s1)
	// 创建一个整型切片
	// 其长度为 3 个元素，容量为 5 个元素
	//slice := make([]int, 3, 5)    //compile error
	s2 := append(slice, "b")
	println(s2)

	// 创建字符串切片
	// 其长度和容量都是 5 个元素
	slice3 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	s3 := append(slice3, "c")
	println(s3)
	// 创建一个整型切片
	// 其长度和容量都是 3 个元素
	slice4 := []int{10, 20, 30}
	s4 := append(slice4, 10)
	println(s4)

	// 创建字符串切片
	// 使用空字符串初始化第 100 个元素
	slice5 := []string{99: "ss"}
	println(slice5)
	println("index 99: " + slice5[99])

	// 创建 nil 整型切片
	var slice6 []int
	println(slice6)

	// 使用 make 创建空的整型切片
	slice7 := make([]int, 0)
	// 使用切片字面量创建空的整型切片
	slice8 := []int{}

	println(slice7)
	println(slice8)

	return s
}

func test2() {
	// 创建一个整型切片
	// 其长度和容量都是 5 个元素
	slice := []int{10, 20, 30, 40, 50}
	// 创建一个新切片
	// 其长度为 2 个元素，容量为 4 个元素
	newSlice := slice[1:3]
	// 修改 newSlice 索引为 3 的元素
	// 这个元素对于 newSlice 来说并不存在
	//newSlice[3] = 45    //runtime error
	println(newSlice)

	// 迭代每一个元素，并显示其值
	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}

	// 创建一个整型切片
	// 其长度和容量都是 4 个元素
	slice2 := []int{10, 20, 30, 40}
	// 迭代每个元素，并显示值和地址
	for index, value := range slice2 {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n",
			value, &value, &slice2[index])
	}

	// 从第三个元素开始迭代每个元素
	for index := 2; index < len(slice); index++ {
		fmt.Printf("Index: %d Value: %d\n", index, slice[index])
	}
}

func test3() {
	// 创建一个整型切片的切片
	slice := [][]int{{10}, {100, 200}}
	// 为第一个切片追加值为 20 的元素
	slice[0] = append(slice[0], 20)
	println(slice)
}

func main() {
	//test("s")
	test3()
}

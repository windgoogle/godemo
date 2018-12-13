package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var m map[int]string = map[int]string{
		90: "Best",
		80: "better",
		70: "good",
	}

	fmt.Println(m, len(m))

	var mp = map[int]string{ //支持类型推导
		90: "优秀",
		80: "良好",
		70: "及格",
	}

	fmt.Println(mp, len(mp))

	var amap = make(map[int]string, 16) //指定容量
	fmt.Println(amap, len(amap))

	//map读写
	var fruits = map[string]int{
		"apple":  8,
		"balana": 5,
		"orange": 8,
	}
	//读取元素
	var score = fruits["apple"]
	fmt.Println(score)

	//增加或修改元素
	fruits["pear"] = 3
	fruits["orange"] = 10
	fmt.Println(fruits)
	//删除元素
	delete(fruits, "pear") //key 不存在时，静默处理，并没有返回值
	fmt.Println(fruits)

	//map 特殊方法

	var sc, ok = fruits["durin"]

	if ok {
		fmt.Println(sc)
	} else {
		fmt.Println("durin not exist")
	}

	fruits["durin"] = 0

	sc, ok = fruits["durin"]

	if ok {
		fmt.Println(sc)
	} else {
		fmt.Println("durin not exist")
	}

	//map 的遍历

	for name, score := range fruits {
		fmt.Println(name, score)
	}

	for name := range fruits {
		fmt.Println(name)
	}

	fmt.Println(unsafe.Sizeof(fruits)) //字典变量存的是指针，指向map头部,所以字典变量内存占用是8字节（64位），32位是4字节

	//内置map 不是线程（协程）安全的，需要用锁来控制

}

package main

import "fmt"

func main() {
	// 创建一个映射，键的类型是 string，值的类型是 int
	dict1 := make(map[string]int)
	// 创建一个映射，键和值的类型都是 string
	// 使用两个键值对初始化映射
	dict2 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}

	println(dict1)
	println(dict2)

	// 创建一个空映射，用来存储颜色以及颜色对应的十六进制代码
	colors := map[string]string{}
	// 将 Red 的代码加入到映射
	colors["Red"] = "#da1337"

	// 通过声明映射创建一个 nil 映射
	//var colors2 map[string]string
	// 将 Red 的代码加入到映射
	//colors2["Red"] = "#da1337"   //panic: assignment to entry in nil map

	// 获取键 Blue 对应的值
	value, exists := colors["Red"] // 这个键存在吗？
	if exists {
		fmt.Println(value)
	}

	// 创建一个映射，存储颜色以及颜色对应的十六进制代码
	colors3 := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	// 显示映射里的所有颜色
	for key, value := range colors3 {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
}

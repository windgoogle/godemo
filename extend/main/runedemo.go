package main

import "fmt"

func main() {
	var s = "计算机科学Compunter Science"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	//codepoint 表示字符串的索引位置，runevalue 表示对应的unicode编码，类型是rune  （北欧战神）
	for codepoint, runevalue := range s {
		fmt.Printf("%d %d ", codepoint, int32(runevalue))
	}

	//字符串是只读的
	//	var s1="hello"fmt.Println();
	//	s1[0]='H'   //cannot assign to s1[0]
	fmt.Println()
	//切割字符串
	var s2 = "hello world"
	var s3 = s2[3:8]

	fmt.Println(s3)

	//字节和字符串的相互转换
	var b = []byte(s2) //字符串转换侧字节切片
	var s4 = string(b) //字节切片转换成字符串

	fmt.Println(b)
	fmt.Println(s4)

	//字符串赋值另一个字符串，底层字符数组是共享的
	//字符串转换成字节数组，不是共享的，因为字符串是只读的，而字节数组不是，因此字符串和字节切片转换，底层数组是需要拷贝的，所以转换操作
	// 存在一定的成本

	//切片的追加，每一次切片追加会形成新的切片变量，如果未扩容，那么追加前后的变量共享原来底层数组，如果

}

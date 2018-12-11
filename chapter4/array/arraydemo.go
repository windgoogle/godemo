package main

func testa() {
	array := [...]int{2: 10, 4: 20}
	array[2] = 5

	// 声明包含 5 个元素的指向整数的数组
	// 用整型指针初始化索引为 0 和 1 的数组元素
	array := [5]*int{0: new(int), 1: new(int)}
	// 为索引为 0 和 1 的元素赋值
	*array[0] = 10
	*array[1] = 20

	// 声明第一个包含 5 个元素的字符串数组
	var array1 [5]string
	// 声明第二个包含 5 个元素的字符串数组
	// 用颜色初始化数组
	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	// 把 array2 的值复制到 array1
	array1 = array2

	// 声明第一个包含 3 个元素的指向字符串的指针数组
	var array3 [3]*string
	// 声明第二个包含 3 个元素的指向字符串的指针数组
	// 使用字符串指针初始化这个数组
	array4 := [3]*string{new(string), new(string), new(string)}
	// 使用颜色为每个元素赋值
	*array4[0] = "Red"
	*array4[1] = "Blue"
	*array4[2] = "Green"
	// 将 array2 复制给 array1
	array1 = array2

	// 声明一个二维整型数组，两个维度分别存储 4 个元素和 2 个元素
	var array5 [4][2]int
	// 使用数组字面量来声明并初始化一个二维整型数组
	array6 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	// 声明并初始化外层数组中索引为 1 个和 3 的元素
	array7 := [4][2]int{1: {20, 21}, 3: {40, 41}}
	// 声明并初始化外层数组和内层数组的单个元素
	array8 := [4][2]int{1: {0: 20}, 3: {1: 41}}
}

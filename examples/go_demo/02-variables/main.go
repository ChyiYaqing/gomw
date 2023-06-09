package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 变量与内置数据类型
	var a int8 = 10
	var a1 byte = 'a'
	var a2 float32 = 12.2
	var a3 = "Hello World"
	a4 := false
	fmt.Println(a, a1, a2, a3, a4)

	// 在Go语言中,字符串使用UTF8编码，UTF8的好处在于，如果基本是英文,每个字符占1byte 和ASCII编码是一样的。非常节省空间。如果是中文，一般占3字节

	str1 := "Golang"
	str2 := "Go语言"
	fmt.Println(reflect.TypeOf(str2[2]).Kind()) // uint8
	fmt.Println(str1[2], string(str1[2]))       // 108 l
	fmt.Printf("%d %c\n", str2[2], str2[2])     // 232 è
	fmt.Println("len(str2): ", len(str2))       // len(str2):  8

	// 转换成[]rune类型后，字符串中的每个字符，无论占多少字节都用int32来表示
	runeArr := []rune(str2)                        // 将string转为rune数组
	fmt.Println(reflect.TypeOf(runeArr[2]).Kind()) // int32
	fmt.Println(runeArr[2], string(runeArr[2]))    // 35821 语
	fmt.Println("len(runeArr): ", len(runeArr))    // len(runeArr):  4

	// 数组array与切片slice
	var arr = [5]int{1, 2, 3, 4, 5} // 一维
	var arr2 [5][5]int              // 二维
	for i := 0; i < len(arr); i++ {
		arr[i] += 100
	}
	fmt.Println(arr)
	fmt.Println(arr2)
	// 切片是数组的抽象, 切片使用数组作为底层结构
	// 切片包含三个组件: 容量，长度，指向底层数组的指针
	slice1 := make([]float32, 0) // 长度为0的切片
	fmt.Println(slice1)
	slice2 := make([]float32, 3, 5) // [0,0,0] 长度为3容量为5的切片
	fmt.Println(len(slice2), cap(slice2))
	// 使用切片，添加元素，切片容量可以根据需要自动扩展
	slice2 = append(slice2, 1, 2, 3, 4)
	fmt.Println(len(slice2), cap(slice2))
	// 子切片 [start, end]
	sub1 := slice2[3:]
	sub2 := slice2[:3]
	sub3 := slice2[1:4]
	fmt.Println(sub1, sub2, sub3)
	// 合并切片
	combined := append(sub1, sub2...) // sub2... 是切片解构的写法，将切片解构为N各独立的元素
	fmt.Println(combined)

	// 字典 map
	m1 := make(map[string]int)
	m1["Tom"] = 18
	// 声明并初始化
	m2 := map[string]string{
		"Sam":   "Male",
		"Alice": "Female",
	}
	fmt.Println(m1, m2)

	// 指针 pointer, 类型定义时使用符号*, 对于一个已经存在的变量，使用&获取该变量的地址
	str := "Golang"
	var p *string = &str // p 是指向str的指针
	*p = "Hello"
	fmt.Println(str)
	// Go语言中，参数是按照值传递的，如果不使用指针，函数内部将会拷贝一份参数的副本，对参数的修改并不会影响到外部变量的值
	// 如果参数使用指针，对参数的传递将会影响到外部变量
	num := 100
	add(num) // num 没有变化
	fmt.Println(num)
	realAdd(&num) // 指针传递，num被修改
	fmt.Println(num)
}

func add(num int) {
	num += 1
}

func realAdd(num *int) {
	*num += 1
}

/*
空值: nil
整型类型: int(取决于操作系统), int8, int16, int32, int64, uint8, uint16
*/

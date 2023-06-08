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

}

/*
空值: nil
整型类型: int(取决于操作系统), int8, int16, int32, int64, uint8, uint16
*/

package main

import "fmt"

// 流程控制
//
//	if
//	for
//	switch
func main() {
	// 条件语句 if else
	age := 18
	if age < 18 {
		fmt.Println("Kid")
	} else {
		fmt.Println("Adult")
	}

	// switch
	gender := MALE
	switch gender {
	case FEMALE:
		fmt.Println("female")
		fallthrough
	case MALE:
		fmt.Println("male")
		fallthrough
	default:
		fmt.Println("unknown")
	}

	// for循环
	sum := 0
	for i := 0; i < 10; i++ {
		if sum > 50 {
			break
		}
		sum += i
	}
	fmt.Println(sum)

	// range 遍历
	nums := []int{10, 20, 30, 40}
	for i, num := range nums {
		fmt.Println(i, num)
	}
}

// 使用type关键字定义一个新的类型Gender
type Gender int8

// 使用const定义MALE,FEMALE两个常量
// Go语言中没有枚举enum类型，可以使用常量的方式模拟枚举
const (
	MALE   Gender = 1
	FEMALE Gender = 2
)

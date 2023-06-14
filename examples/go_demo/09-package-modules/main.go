package main

import "fmt"

// 一个文件夹可以作为package,同一个package内部变量、类型、方法等定义可以相互看到
// Go语言有Public和Private的概念，粒度是包，如果类型/接口/方法/函数/字段的首字母大写，则是Public的，对其他package可见

func main() {
	fmt.Println(add(3, 5))
}

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("filename.txt")
	if err != nil {
		fmt.Println(err)
	}
	if err = hello(""); err != nil {
		fmt.Println(err)
	}
	fmt.Println(get(5))
	fmt.Println("finished")
}

func hello(name string) error {
	if len(name) == 0 {
		// errors.New 返回自定义的错误
		return errors.New("error: name is null")
	}
	fmt.Println("Hello, ", name)
	return nil
}

// Go语言提供的defer和recover
func get(index int) (ret int) {
	// 使用defer 定义异常处理的函数, 在协程退出前，会执行完defer挂载的任务
	defer func() {
		// 使用recover 使的程序恢复正常
		if r := recover(); r != nil {
			fmt.Println("Some error happend!", r)
			ret = -1
		}
	}()
	arr := [3]int{2, 3, 4}
	return arr[index]
}

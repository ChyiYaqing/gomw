package main

import "fmt"

func add(num1 int, num2 int) (ans int) {
	ans = num1 + num2
	return ans
}

func div(num1 int, num2 int) (int, int) {
	return num1 / num2, num1 % num2
}

func main() {
	quo, rem := div(100, 17)
	fmt.Println(quo, rem)
	fmt.Println(add(100, 17))
}

package gomonkey_demo

// 定义一个加法方法
func Add(a, b int) int {
	return a + b
}

// 定义方法，返回一个整数的2倍
func GetDouble(a int) int {
	return Add(a, a)
}

package unit_demo

import "fmt"

func Hello() string {
	return "Hello, world"
}

func Add(a, b int) int {
	return a + b
}

// 匹配URL
func CheckUrl(url string) bool {
	var urlList = [2]string{"learnku.com", "xdcute.com"}
	for v := range urlList {
		if urlList[v] == url {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(Hello())
	Add(1, 2)
}

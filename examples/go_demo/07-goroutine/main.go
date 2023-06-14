package main

import (
	"fmt"
	"sync"
	"time"
)

// Go预演提供sync和channel两种方式支持协程goroutine的并发
func main() {

	// sync.WaitGroup 等待并发协程执行
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go download("a.com/"+string(i+'0'), &wg)
	}
	wg.Wait() // 等待所有的协程执行结束
	fmt.Println("Done!")

	var ch = make(chan string, 10) // 创建大小为 10的缓冲信道
	for i := 0; i < 3; i++ {
		go download2("a.com/"+string(i+'0'), ch)
	}
	for i := 0; i < 3; i++ {
		msg := <-ch // 等待信道返回消息
		fmt.Println("finish", msg)
	}
	fmt.Println("Done!")
}

func download(url string, wg *sync.WaitGroup) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second) // 模拟耗时操作
	wg.Done()
}

func download2(url string, ch chan string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	ch <- url // 将url发送给信道
}

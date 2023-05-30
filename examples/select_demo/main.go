package main

import (
	"fmt"
	"time"
)

type Msg struct {
	text     string
	lastTime time.Time
}

// 超时推出逻辑
func main() {
	ch := make(chan int, 3)
	timeoutCh := make(chan bool)
	msgCh := make(chan Msg)
	go func(ch chan int) {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			ch <- 1
		}
	}(ch)
	go func(timeoutCh chan bool, msgCh chan Msg) {
		lasttime := time.Now()
		for {
			select {
			case msg := <-msgCh:
				lasttime = msg.lastTime
			default:
				if time.Now().Sub(lasttime).Seconds() >= 10 {
					timeoutCh <- true
					return
				}
				time.Sleep(time.Second)
			}
		}
	}(timeoutCh, msgCh)
	for {
		select {
		case <-timeoutCh:
			fmt.Println("超时退出", time.Now().Format(time.RFC3339))
			return
		case v := <-ch:
			msgCh <- Msg{
				text:     "text",
				lastTime: time.Now(),
			}
			fmt.Printf("收到新数据v: %d, time: %s \n", v, time.Now().Format(time.RFC3339))
		}
	}
}

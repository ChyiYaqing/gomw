package main

import (
	"fmt"
	"time"
)

func main() {

	startTime := time.Now()
	time.Sleep(10 * time.Second)
	endTime := time.Now()

	if endTime.Sub(startTime).Seconds() > 20 {
		fmt.Println("timeout")
	} else {
		fmt.Println("no timeout")
	}
}

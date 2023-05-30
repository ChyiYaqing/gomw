package main

import (
	"fmt"
	"time"
)

/*
	watchdog -- A simple watchdog timer based on golang channels




*/

// Watchdog ... use the GetKickChannel() to get notified when the watchdog barks.
type Watchdog struct {
	timer    *time.Timer
	timeout  time.Duration
	kickChan chan bool
	stopChan chan bool
}

// NewWatchdog ... ctor
func NewWatchdog(timeout time.Duration) *Watchdog {
	w := &Watchdog{
		timer:    time.NewTimer(timeout),
		timeout:  timeout,
		kickChan: make(chan bool, 1),
		stopChan: make(chan bool, 1),
	}
	w.timerExpireWorker()
	return w
}

func (w *Watchdog) timerExpireWorker() {
	go func(w *Watchdog) {
		for {
			select {
			case <-w.stopChan:
				return
			case <-w.timer.C:
				w.kickChan <- true
				return
			}
		}
	}(w)
}

// GetKickChannel ... returns the kick channel
func (w *Watchdog) GetKickChannel() chan bool {
	return w.kickChan
}

// Stop ... stops the watchdog
func (w *Watchdog) Stop() {
	w.timer.Stop()
	w.stopChan <- true
}

// Reset ... resets the timer
func (w *Watchdog) Reset() {
	w.timer.Stop()
	w.timer = time.NewTimer(w.timeout)
}

func main() {
	ch := make(chan int)
	go func(ch chan int) {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Microsecond * 10)
		}
		return
	}(ch)
	// kicks at interval on the GetKickChannel()
	w := NewWatchdog(time.Second * 3)
	for {
		select {
		case <-w.GetKickChannel():
			fmt.Println("timeout kick!")
			return
		case v := <-ch:
			fmt.Println(v)
			w.Reset() // 喂狗
			time.Sleep(time.Second)
		}
	}
}

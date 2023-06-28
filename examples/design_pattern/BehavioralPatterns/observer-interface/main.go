package main

import "fmt"

/*
the Observer Pattern: 观察者模式
	> allows an object (known as the subject) to notify its dependent objects (known as observers) automatically
	when a change occurs in the subject.

	In Golang, there are two main ways to implement the Observer Pattern:
		1. using interfaces
		2. using channels
*/

// the classic Observer Pattern -- using interface
// the subject maintains a list of observers, and provides methods to register and remove observers from the list.

type subject interface {
	registerObserver(observer) // 注册
	removeObserver(observer)   // 移除
	notifyObservers()          // 通知
}

type observer interface {
	update()
}

type weatherStation struct {
	temperature float64
	observers   []observer
}

type temperatureDisplay struct {
	temperature float64
}

func (w *weatherStation) registerObserver(o observer) {
	w.observers = append(w.observers, o)
}

func (w *weatherStation) removeObserver(o observer) {
	for i, observer := range w.observers {
		if observer == o {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			break
		}
	}
}

func (w *weatherStation) notifyObservers() {
	for _, observer := range w.observers {
		observer.update()
	}
}

func (t *temperatureDisplay) update() {
	fmt.Printf("Temperature: %.2f\n", t.temperature)
}

func main() {

}

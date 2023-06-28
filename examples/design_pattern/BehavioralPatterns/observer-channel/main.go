package main

/*
the Observer Pattern in Golang using channels
*/

type subject struct {
	observers []chan float64
}

type observer struct {
	c chan float64
}

func (s *subject) registerObserver() chan float64 {
	c := make(chan float64)
	s.observers = append(s.observers, c)
	return c
}

func (s *subject) removeObserver(c chan float64) {
	for i, observer := range s.observers {
		if observer == c {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *subject) notifyObservers(temperature float64) {
	for _, observer := range s.observers {
		go func(c chan float64) {
			c <- temperature
		}(observer)
	}
}

func main() {

}

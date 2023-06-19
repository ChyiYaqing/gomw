package main

import "sync"

/*
Singleton Pattern
	>
*/

type Singleton struct {
	name string
}

var (
	instance *Singleton
	once     sync.Once
)

// ensure that only one instance of the Singleton struct is created, and that the creation is thread-safe.
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{name: "Safe Golang Singleton"}
	})
	return instance
}

func (s *Singleton) GetName() string {
	return s.name
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()
	println("s1.GetName()", s1.GetName())
	println("s2.GetName()", s2.GetName())
	println("s1 == s2", s1 == s2)

}

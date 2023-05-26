package main

import (
	"fmt"
	"time"

	"golang.org/x/sync/singleflight"
)

// singleflight prevents similar function invocations at the same time.
// This is done by allowing the second invocation to wait for the
// first invocation to complete. Once complete, the result of the first invocation
// is shared with the second.

func main() {
	case1()
}

// Use case 1 - net/lookup
func case1() {
	var g singleflight.Group
	go func() {
		// func (g *Group) Do(key string, fn func() (interface{}, error))
		v, err, shared := g.Do("key", func() (interface{}, error) {
			s := timeConsumingFunction()
			return s, nil
		})
		fmt.Printf("Goroutine result : %v, err: %v, shared: %v\n", v, err, shared)
	}()
	time.Sleep(2 * time.Second)

	v, err, shared := g.Do("key", func() (interface{}, error) {
		s := timeConsumingFunction()
		return s, nil
	})
	fmt.Printf("Goroutine result : %v, err: %v, shared: %v\n", v, err, shared)
}

func timeConsumingFunction() string {
	time.Sleep(10 * time.Second)
	fmt.Println("exe Hello")
	return "Hello"
}

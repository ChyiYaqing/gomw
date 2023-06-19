package main

import "fmt"

/*

Prototype Pattern
	1. define an interface that includes a method to clone the object.
	2. define a struct that implements this interface and provides a clone method that returns a copy of itself.

the Prototype pattern can be a useful tool for creating objects in a more efficient and flexible way.
*/

type Cloneable interface {
	Clone() Cloneable
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) Clone() Cloneable {
	return &Person{
		Name: p.Name,
		Age:  p.Age,
	}
}

func main() {

	person1 := &Person{Name: "Alice", Age: 30}
	person2 := person1.Clone().(*Person)
	fmt.Println(person1)
	fmt.Println(person2)
}

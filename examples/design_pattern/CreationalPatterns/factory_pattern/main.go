package main

import "fmt"

/*
the Factory Method Pattern
	> Instead of creating objects directly, we delegate the creation to a factory that can
	decide which class of object to create based on certain conditions.

the Advantage of Using the Factory Method Pattern
	1. Flexibility: The Factory Method pattern allows us to create objects without knowing their exact class.
	2. Scalability: The Factory Method pattern can be easily scaled to support new types of objects.
	3. Decoupling: The Factory Method pattern decouples the creation of objects from their usage.
*/

type Product interface {
	GetName() string
}

type Factory interface {
	CreateProduct() Product
}

type ConcreteProduct struct {
	name string
}

func (p *ConcreteProduct) GetName() string {
	return p.name
}

type ConcreteFactory struct{}

// return Product interface
func (f *ConcreteFactory) CreateProduct() Product {
	return &ConcreteProduct{name: "Concrete Product"}
}

func main() {
	factory := &ConcreteFactory{}
	product := factory.CreateProduct()
	fmt.Println(product.GetName())
}

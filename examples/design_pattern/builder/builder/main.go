package main

import "fmt"

/*
the Builder pattern in Golang
	1. A builder interface that defines the steps for constructing the object.
	2. A ConcreteBuilder struct that implements the Builder interface and provides a way to construct the object.
	3. A Direct struct that uses the Builder to construct the object.
*/

type Pizza struct {
	dough    string
	sauce    string
	cheese   string
	toppings []string
}

type PizzaBuilder interface {
	SetDough(string) PizzaBuilder
	SetSauce(string) PizzaBuilder
	SetCheese(string) PizzaBuilder
	SetToppings([]string) PizzaBuilder
	Build() *Pizza
}

type ConcreatePizzaBuilder struct {
	pizza *Pizza
}

func NewConcretePizzaBuilder() *ConcreatePizzaBuilder {
	return &ConcreatePizzaBuilder{pizza: &Pizza{}}
}

func (cpb *ConcreatePizzaBuilder) SetDough(dough string) PizzaBuilder {
	cpb.pizza.dough = dough
	return cpb
}

func (cpb *ConcreatePizzaBuilder) SetSauce(sauce string) PizzaBuilder {
	cpb.pizza.sauce = sauce
	return cpb
}

func (cpb *ConcreatePizzaBuilder) SetCheese(cheese string) PizzaBuilder {
	cpb.pizza.cheese = cheese
	return cpb
}

func (cpb *ConcreatePizzaBuilder) SetToppings(toppings []string) PizzaBuilder {
	cpb.pizza.toppings = toppings
	return cpb
}

func (cpb *ConcreatePizzaBuilder) Build() *Pizza {
	return cpb.pizza
}

type Director struct {
	builder PizzaBuilder
}

func NewDirector(builder PizzaBuilder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct() *Pizza {
	return d.builder.SetDough("Thin Crust").SetSauce("Tomato").SetCheese("Mozzarella").SetToppings([]string{"Mushrooms", "Olives", "Onions"}).Build()
}

func main() {
	builder := NewConcretePizzaBuilder()
	director := NewDirector(builder)
	pizza := director.Construct()
	fmt.Println(pizza)
}

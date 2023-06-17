package main

/*
	the Options pattern:


	the Builder pattern
*/

// The Options Pattern
type Pizza struct {
	dough    string
	sauce    string
	cheese   string
	toppings []string
}

type PizzaOptions struct {
	Dough    string
	Sauce    string
	Cheese   string
	Toppings []string
}

type PizzaOption func(*PizzaOptions)

func WithDough(dough string) PizzaOption {
	return func(po *PizzaOptions) {
		po.Dough = dough
	}
}

func WithSauce(sauce string) PizzaOption {
	return func(po *PizzaOptions) {
		po.Sauce = sauce
	}
}

func WithCheese(cheese string) PizzaOption {
	return func(po *PizzaOptions) {
		po.Cheese = cheese
	}
}

func WithToppings(toppings []string) PizzaOption {
	return func(po *PizzaOptions) {
		po.Toppings = toppings
	}
}

func NewPizza(options ...PizzaOption) *Pizza {
	opts := &PizzaOptions{}
	for _, option := range options {
		option(opts)
	}
	pizza := &Pizza{
		dough:    opts.Dough,
		sauce:    opts.Sauce,
		cheese:   opts.Cheese,
		toppings: opts.Toppings,
	}
	return pizza
}

func main() {
	pizza := NewPizza(
		WithDough("Regular"),
		WithSauce("Tomato"),
		WithCheese("Mozzarella"),
		WithToppings([]string{"Pepperoni", "Olives", "Mushrooms"}),
	)
	println(pizza.dough)
	println(pizza.sauce)
	println(pizza.cheese)
	println(pizza.toppings)

}

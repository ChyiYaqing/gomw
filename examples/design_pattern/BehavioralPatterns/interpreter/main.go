package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
the Interpreter Pattern: 解释者模式


*/

type Product struct {
	Name  string
	Price int
	Brand string
	Color string
}

type Context struct {
	Products []Product
}

type Interpreter interface {
	Interpret(context *Context) []Product
}

type Parser struct {
	stack []Interpreter
}

func (p *Parser) Parse(expression string) {
	tokens := strings.Split(expression, " ")
	for _, token := range tokens {
		parts := strings.Split(token, ":")
		switch parts[0] {
		case "price":
			args := strings.Trim(parts[1], "range")
			prices := strings.Split(args, ",")
			min, _ := strconv.Atoi(prices[0])
			max, _ := strconv.Atoi(prices[1])
			filter := &PriceFilter{min: min, max: max}
			p.stack = append(p.stack, filter)
		case "brand":
			brand := parts[1]
			filter := &BrandFilter{brand: brand}
			p.stack = append(p.stack, filter)
		case "color":
			color := parts[1]
			filter := &ColorFilter{color: color}
			p.stack = append(p.stack, filter)
		}
	}
}

func (p *Parser) Result(context *Context) []Product {
	for _, i := range p.stack {
		context.Products = i.Interpret(context)
	}
	return context.Products
}

type PriceFilter struct {
	min int
	max int
}

func (pf *PriceFilter) Interpret(context *Context) []Product {
	products := context.Products
	filtered := []Product{}
	for _, p := range products {
		if p.Price >= pf.min && p.Price <= pf.max {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

type BrandFilter struct {
	brand string
}

func (bf *BrandFilter) Interpret(context *Context) []Product {
	products := context.Products
	filtered := []Product{}
	for _, p := range products {
		if p.Brand == bf.brand {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

type ColorFilter struct {
	color string
}

func (cf *ColorFilter) Interpret(context *Context) []Product {
	products := context.Products
	filtered := []Product{}
	for _, p := range products {
		if p.Color == cf.color {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

func main() {
	products := []Product{
		{Name: "Product 1", Price: 100, Brand: "Brand", Color: "Red"},
		{Name: "Product 2", Price: 50, Brand: "Brand", Color: "Blue"},
		{Name: "Product 3", Price: 200, Brand: "Brand", Color: "Green"},
	}
	context := &Context{Products: products}
	parser := &Parser{}
	parser.Parse("price:range(50, 250) brand:Brand color:Green")
	filtered := parser.Result(context)
	fmt.Println(filtered)
}

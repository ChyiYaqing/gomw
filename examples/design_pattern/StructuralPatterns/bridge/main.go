package main

import "fmt"

/*
the Bridge Pattern

*/

type Renderer interface {
	RenderCircle(radius float32)
	RenderSquare(side float32)
}

type Shape interface {
	Draw()
}

type Circle struct {
	x, y, radius float32
	renderer     Renderer
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

type Square struct {
	side     float32
	renderer Renderer
}

func (s *Square) Draw() {
	s.renderer.RenderSquare(s.side)
}

type VectorRenderer struct{}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Printf("Drawing a circle of radius %f in VectorRenderer\\\\n", radius)
}

func (v *VectorRenderer) RenderSquare(side float32) {
	fmt.Printf("Drawing a square of side %f in VectorRenderer\\\\n", side)
}

type RasterRenderer struct{}

func (r *RasterRenderer) RenderCircle(radius float32) {
	fmt.Printf("Drawing a circle of radius %f in RasterRenderer\\\\n", radius)
}

func (r *RasterRenderer) RenderSquare(side float32) {
	fmt.Printf("Drawing a square of side %f in RasterRenderer\\\\n", side)
}

func main() {
	circle := &Circle{x: 0, y: 0, radius: 5, renderer: &VectorRenderer{}}
	square := &Square{side: 10, renderer: &RasterRenderer{}}

	circle.Draw() // Drawing a circle of radius 5.000000 in VectorRenderer
	square.Draw() // Drawing a square of side 10.000000 in RasterRenderer
}

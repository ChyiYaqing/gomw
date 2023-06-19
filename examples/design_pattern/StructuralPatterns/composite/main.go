package main

import "fmt"

/*
the Composite Pattern: 复合模式
	> The Composite pattern is a powerful tool for working with hierarchical structures.
	> It allows you to create a tree-like structure which each node in the tree can be either a single object or a group of objects.

*/

type Component interface {
	// The Operation method is the method that will be called on all the objects in the tree.
	Operation()
}

type Composite struct {
	children []Component
}

func (c *Composite) Operation() {
	fmt.Println("Composite operation")
	for _, child := range c.children {
		child.Operation()
	}
}

type Leaf struct {
	data string
}

func (l *Leaf) Operation() {
	fmt.Println("Leaf operation -", l.data)
}

func main() {
	root := &Composite{
		children: []Component{
			&Leaf{data: "file1.txt"},
			&Composite{
				children: []Component{
					&Leaf{data: "file2.txt"},
					&Composite{
						children: []Component{
							&Leaf{data: "file3.txt"},
						},
					},
				},
			},
		},
	}
	root.Operation()
}

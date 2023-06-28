package main

import "fmt"

/*
the	Iterator Pattern: 迭代器模式
	> The Ierator Pattern provides a way to access the elements of an aggregate
	object sequentially without exposing its underlying implementation.

*/

type Iterator interface {
	Next() interface{} // returns the next element in the collection
	HasNext() bool     // checks whether there are more elements to traverse.
}

// The Slice[T] type is a generic slice.
type Slice[T any] []T

func (s Slice[T]) Iterator() Iterator[T] {
	return &sliceIterator{T, slice: s, index: -1}
}

type sliceIterator[T any] struct {
	slice Slice[T]
	index int
}

func (i *sliceIterator[T]) Next() T {
	i.index++
	return i.slice[i.index]
}

func (i *sliceIterator[T]) HasNext() bool {
	return i.index+1 < len(i.slice)
}

func main() {
	s := Slice[int]{1, 2, 3, 4, 5}
	it := s.Iterater()
	for it.HasNext() {
		fmt.Println(it.Next())
	}

}

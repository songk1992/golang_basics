package lec017

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	height float64
}

func (rect Rect) Area() float64 {
	return rect.height * rect.width
}

func (rect Rect) Perimeter() float64 {
	return 2 * (rect.height + rect.width)
}

func Lec017() {
	var shape Shape
	shape = Rect{3.0, 4.0}
	rect := Rect{3.0, 4.0}
	fmt.Printf("%T %v %v %v %v", shape, shape, shape.Area(), shape.Perimeter(), rect == shape)
}
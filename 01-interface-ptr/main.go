package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Circumference() float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Circumference() float64 {
	return 2 * (r.Length + r.Width)
}

// a struct example where one of the receiver function expects a pointer
type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return math.Pow(s.Side, 2)
}

func (s *Square) Circumference() float64 {
	return 4 * s.Side
}

func main() {
	var s Shape
	r := Rectangle{10.0, 5.0}
	s = r

	fmt.Println(s == r) // this will return true

	// this will also return true. looks like it checks by value
	s = Rectangle{10.0, 5.0}
	fmt.Println(s == r, s == Rectangle{10.0, 5.0})

	var s2 Shape
	// however, this one results in compilation error, since Circumference is a pointer receiver
	// s2 = Square{10}
	// this one is valid. Somehow non-pointer receiver still valid?
	sq := &Square{10}
	s2 = sq
	fmt.Println(s2.Area(), s2.Circumference())
	// since this line below compares pointers instead of value, the former will return false
	// and the latter will return true
	fmt.Println(s2 == sq, s2 == &Square{10})
}

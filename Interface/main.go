package main

import (
	"fmt"
	"math"
	"reflect"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius int64
}

type Square struct {
	Side int64
}

type Rectangle struct {
	Length  int64
	Breadth int64
}

func (c Circle) Area() float64 {
	radius := c.Radius
	return math.Pi * float64(radius) * float64(radius)
}

func (c Circle) Perimeter() float64 {
	r := c.Radius
	return 2 * math.Pi * float64(r)
}

func (s Square) Perimeter() float64 {
	si := s.Side
	return 4 * float64(si)
}

func (s Square) Area() float64 {
	si := s.Side
	return float64(si) * float64(si)
}

func (r Rectangle) Area() float64 {
	le := r.Length
	br := r.Breadth
	return float64(le) * float64(br)
}

func (r Rectangle) Perimeter() float64 {
	le := r.Length
	br := r.Breadth
	return 2 * (float64(le) + float64(br))
}

func calculateDimension(s Shape) {
	ty := reflect.TypeOf(s)
	str := ty.String()
	sh := str[5:]

	fmt.Println(fmt.Sprintf("The Area of %s is %f", sh, s.Area()))
	fmt.Println(fmt.Sprintf("The Perimeter of %s is %f", sh, s.Perimeter()))

}

func main() {
	ci := Circle{
		Radius: 10,
	}

	re := Rectangle{
		Length:  10,
		Breadth: 8,
	}

	sq := Square{
		Side: 10,
	}

	calculateDimension(ci)
	calculateDimension(sq)
	calculateDimension(re)
}

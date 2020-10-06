package structs

import "math"

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Length + r.Width)
}

func Area(r Rectangle) float64 {
	return r.Length * r.Width
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Length, Width float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

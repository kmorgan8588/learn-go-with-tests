package structs

import "math"

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Length + r.Width)
}

func Area(r Rectangle) float64 {
	return r.Length * r.Width
}

type Rectangle struct {
	Length, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

type Shape interface {
	Area() float64
}

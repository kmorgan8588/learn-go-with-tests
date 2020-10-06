package structs

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Length + r.Width)
}

func Area(r Rectangle) float64 {
	return r.Length * r.Width
}

type Rectangle struct {
	Length, Width float64
}

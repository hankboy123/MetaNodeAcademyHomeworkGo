package main

type Rectangle struct {
	Length float64
	Width  float64
}

func (c Rectangle) Area() float64 {
	return c.Length * c.Width
}

func (c Rectangle) Perimeter() float64 {
	return 2 * (c.Length + c.Width)
}

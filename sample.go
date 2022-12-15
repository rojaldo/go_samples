package main

type Square struct {
	side int
}

type Rectangle struct {
	side1 int
	side2 int
}

type Circle struct {
	radius int
}

type Shape interface {
	Area() float32
}

func (s Square) Area() float32 {
	return float32(s.side * s.side)
}

func (r Rectangle) Area() float32 {
	return float32(r.side1 * r.side2)
}

func (c Circle) Area() float32 {
	return float32(c.radius*c.radius) * 3.14
}

package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	side float64
}

func (s square) getArea() float64 {
	return s.side * s.side
}

type triangle struct {
	base  float64
	heigh float64
}

func (t triangle) getArea() float64 {
	return (t.base * t.heigh) / 2
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	t := triangle{base: 3, heigh: 4}
	s := square{side: 4}

	printArea(s)
	printArea(t)
}

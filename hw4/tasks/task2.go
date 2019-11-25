package tasks

import (
	"fmt"
	"math"
)

// Figure consists of needful methods
type Figure interface {
	Area() float64
	Perimeter() float64
}

// Square consists of side of a square
type Square struct {
	A float64
}

// Area calculates area of square
func (s Square) Area() float64 {
	return s.A * s.A
}

// Perimeter calculates perimeter of square
func (s Square) Perimeter() float64 {
	return 4 * s.A
}

// Circle consists of radius of circle
type Circle struct {
	Radius float64
}

// Area calculates area of circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter calculates perimeter of circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Task2 : Implement Square and Circle structures which implements Figure interface.
func Task2() {
	var s Figure = Square{2}
	var c Figure = Circle{18}

	fmt.Println(s.Area(), s.Perimeter())
	fmt.Println(c.Area(), c.Perimeter())

}

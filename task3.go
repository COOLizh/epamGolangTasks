package main

import "fmt"

// Point indicates the coordinates of a point on a plane
type Point struct {
	x, y int
}

// Square contains the necessary data to describe the square
type Square struct {
	start Point
	a     uint
}

// End сalculates the end coordinates of a square
func (s Square) End() (Point) {
	var end Point = Point{s.start.x + int(s.a), s.start.y + int(s.a)}
	return end 
}

// Perimeter calculates the perimeter of a square
func (s Square) Perimeter() uint {
	return 4 * s.a
}

// Area calculates the area of the square
func (s Square) Area() uint {
	return s.a * s.a
}

func main() {
	s := Square{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}

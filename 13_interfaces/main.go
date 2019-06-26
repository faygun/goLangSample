package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	circumference() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) circumference() float64 {
	return 2.0 * math.Pi * c.radius
}

func (r Rectangle) circumference() float64 {
	return (r.width + r.height) * 2.0
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func getCircum(s Shape) float64 {
	return s.circumference()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 18, height: 5}

	fmt.Printf("Circle area : %f\n", getArea(circle))
	fmt.Printf("Rectangle area : %f\n", getArea(rectangle))

	fmt.Printf("Circle circum : %f\n", getCircum(circle))
	fmt.Printf("Rectangle circum : %f\n", getCircum(rectangle))
}

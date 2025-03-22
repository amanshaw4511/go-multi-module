package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (r rect) diag() float64 {
	return math.Sqrt(r.width*r.width + r.height*r.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s Shape) {
	fmt.Println("Area: ", s.area())
	fmt.Println("Perimeter: ", s.perimeter())
}

func interfaces() {
	fmt.Println("\n# Interfaces")

	// Interface defines the behavior of a type instead of the structure
	// It is collection of method signatures
	// A type implements an interface if it has all the methods defined in the interface
	// Interfaces are implemented implicitly in Go, no implements keyword like other languages required

	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	printShapeInfo(r)
	printShapeInfo(c)

	// type casting and type assertion
	var s Shape = rect{width: 3, height: 4}

	sr, ok := s.(rect)
	if !ok {
		fmt.Println("Conversion failed")
	} else {
		fmt.Println("Conversion successful, diagonal: ", sr.diag())
	}

	// type switch
	switch rr := s.(type) {
	case rect:
		fmt.Printf("Reactangle width: %f, height: %f\n", rr.width, rr.height)
	case circle:
		fmt.Println("Circle radius: ", rr.radius)
	default:
		fmt.Println("Unknown")
	}

}

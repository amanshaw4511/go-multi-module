package main

import (
	"fmt"
	"math"
)

// simple struct
type point struct {
	x int
	y int
}

func (p point) isOrigin() bool {
	return p.x == 0 && p.y == 0
}

func (p point) magnitudeAndDirection() (float64, float64) {
	// calculate magnitude and direction of the point
	magnitute := math.Sqrt(float64(p.x*p.x + p.y*p.y))
	direction := math.Atan2(float64(p.y), float64(p.x))
	return magnitute, direction
}

type line struct {
	start point
	end   point
}

// nested struct
type person struct {
	name string
	age  int
	addr struct {
		city string
		zip  int
	}
}

// embedded struct
type indian struct {
	person
	aadharNo string
}

func structs() {
	fmt.Println("\n# Structs")

	// initializing a struct
	p1 := point{10, 20}
	fmt.Println(p1)

	l1 := line{p1, point{30, 40}}
	fmt.Println(l1)

	// nested struct
	John := person{
		name: "John",
		age:  27,
		addr: struct {
			city string
			zip  int
		}{
			city: "New York",
			zip:  202023,
		},
	}
	fmt.Println(John)

	// accessing struct fields
	fmt.Println("Name:", John.name)
	fmt.Println("Age:", John.age)
	fmt.Println("City:", John.addr.city)
	// modifying struct fields
	John.age = 28
	fmt.Println("Modified Age:", John.age)

	// creating struct with zero(default) values
	anonymous := person{}
	fmt.Printf("Anonymous: %+v\n", anonymous)

	// anonymous struct
	ano := struct {
		name string
		age  int
	}{
		name: "ano",
		age:  30,
	}
	fmt.Printf("ano: %+v\n", ano)

	// embedded struct
	ram := indian{
		person: person{
			name: "Ram",
			age:  30,
		},
		aadharNo: "1234567890",
	}
	fmt.Println("age:", ram.age)
	fmt.Println("city:", ram.addr.city)
	ram.age = 99
	fmt.Println("Modified age age:", ram.age)

	fmt.Printf("Ram: %+v\n", ram)

	// struct methods
	point1 := point{3, 4}
	mag, dir := point1.magnitudeAndDirection()
	fmt.Println("Magnitude:", mag)
	fmt.Println("Direction:", dir)
	fmt.Println("Is Origin:", point1.isOrigin())

	// New Concepts
	// 1. embedded struct
	// 2. struct methods
}

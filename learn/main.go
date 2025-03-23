package main

import (
	"fmt"
)

func main() {
	variables()
	basic_types()
	constants()
	comments()
	formatting_strings()
	if_else()
	switch_case()
	loops()
	functions()
	structs()
	interfaces()
	slices()
}

func variables() {
	fmt.Println("\n# Variables")

	var x int  // declaring a variable
	x = 10     // initializing a variable
	var y = 20 // declaring and initializing a variable
	z := 30    // shorthand for declaring and initializing a variable
	fmt.Println(x+y == z)
}

func constants() {
	fmt.Println("\n# Constants")

	// Constants are declared like variables, but with the const keyword.
	// Constants cannot be declared using the := syntax.
	// Constants can be primitive types like strings, integers, booleans and floats. They can not be more complex types like slices, maps and struct
	// Constants are evaluated at compile time.
	// As the name implies, the value of a constant can't be changed after it has been declared.
	// Constants can be shadowed in inner scopes.
	const a = 10
	const b = 20
	const c = a + b // computed constant
	// const d = os.Args[0] // Won't work as os.Args[0] is runtime value
	fmt.Println(c)
}

func basic_types() {
	fmt.Println("\n# Basic Types")

	// Basic types
	// bool
	// string
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte // alias for uint8
	// rune // alias for int32
	// float32 float64
	// complex64 complex128

	// int and uint are machine dependent either 32 or 64 bits

	// Example
	var a bool = true
	var b string = "Hello, World!"
	var c int = -10
	var d uint = 10
	var e float32 = 10.5
	var f complex64 = 10 + 5i
	fmt.Println(a, b, c, d, e, f)

	// Type conversion
	var g float32 = 9.8
	var h int = int(g)
	fmt.Println(h)
}

func comments() {
	fmt.Println("\n# Comments")

	// This is a single line comment
	/*
		This is a multi-line comment
		This is a multi-line comment 2
	*/
}

func formatting_strings() {
	fmt.Println("\n# Formatting Strings")

	var a = 10
	var b = 20

	point := struct {
		x int
		y int
	}{
		x: 10,
		y: 20,
	}

	fmt.Printf("a: %v, b: %v\n", a, b) // %v prints the value in a default format
	fmt.Printf("a: %T, b: %T\n", a, b) // %T prints the type of the variable
	fmt.Printf("point: %+v\n", point)  // %+v adds field names useful for structs

	name := "John"
	age := 27
	greet := fmt.Sprintf("My name is %v and I am %v years old", name, age) // Sprintf returns the formatted string
	fmt.Println(greet)
}

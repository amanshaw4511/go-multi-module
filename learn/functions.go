package main

import "fmt"

func sayHello() {
	fmt.Println("Hello from the functions module!")
}

func greet(name string) {
	fmt.Println("Hello, " + name)
}

func divmod(a, b int) (int, int) {
	return a / b, a % b
}

func divmodNamed(a, b int) (div int, rem int) {
	div = a / b
	rem = a % b
	return // no need to specify return values
}

func evaluate(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a, b int) int {
	return a + b
}

func deferExample() {
	fmt.Println("Defer func Start")
	defer fmt.Println("Defer")
	defer fmt.Println("Defer 2")
	fmt.Println("Defer func End")
}

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func functions() {
	fmt.Println("\n# Functions")

	// function with no arguments
	sayHello()

	// function with arguments
	greet("John")

	// function with multiple return values
	div, rem := divmod(5, 2)
	fmt.Println("divmod(5, 2) =>", div, rem)

	// function with named return values
	div, rem = divmodNamed(5, 2)
	fmt.Println("divmodNamed(5, 2) =>", div, rem)

	// ignore return values
	div, _ = divmod(5, 2)
	fmt.Println("divmod(5, 2) =>", div)

	// function as argument
	result := evaluate(5, 2, add)
	fmt.Println("evaluate(5, 2, add) =>", result)

	// anonymous function
	subtract := func(a, b int) int {
		return a - b
	}
	result = evaluate(5, 2, subtract)
	fmt.Println("evaluate(5, 2, subtract) =>", result)

	// defer keyword
	// defer function is executed after the surrounding function returns
	// in case of multiple defer functions, they are executed in LIFO order
	// defer is useful to close resources like files, connections, etc.
	deferExample()

	// closures
	// a closure is a function value that references variables from outside its body
	// the function may access and assign to the referenced variables
	nextCount := counter()
	fmt.Println(nextCount()) // 1
	fmt.Println(nextCount()) // 2
	fmt.Println(nextCount()) // 3

}

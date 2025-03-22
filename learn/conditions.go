package main

import "fmt"

func if_else() {
	fmt.Println("\n# if_else")

	age := 18

	// if
	if age >= 18 {
		println("Adult")
	}

	// if else
	if age >= 18 {
		println("Adult")
	} else {
		println("Minor")
	}

	// if else if
	if age >= 18 {
		println("Adult")
	} else if age >= 13 {
		println("Teen")
	} else {
		println("Child")
	}

	// if with a initial statement
	// age is only available in the if-else block scope
	if age := 17; age >= 18 {
		println("Adult")
	} else if age >= 13 {
		println("Teen")
	} else {
		println("Child")
	}

	// Note: No ternary operator in Go
}

func switch_case() {
	fmt.Println("\n# switch_case")

	age := 18

	category := ""

	switch {
	case age >= 18:
		category = "Adult"
	case age >= 13:
		category = "Teen"
	default:
		category = "Child"
	}

	fmt.Println("Category:", category)

	switch category {
	case "Adult":
		fmt.Println("Age is 18 or more")
	case "Teen":
		fmt.Println("Age is between 13 and 18")
	default:
		fmt.Println("Age is less than 13")
	}

	// switch with a initial statement
	switch category := "Adult"; category {
	case "Adult":
		fmt.Println("Age is 18 or more")
	case "Teen":
		fmt.Println("Age is between 13 and 18")
	default:
		fmt.Println("Age is less than 13")
	}

	// Notice that in Go, the break statement is not required at the end of a case
	//  to stop it from falling through to the next case. The break statement is implicit in Go.

	// fallthrough keyword is used to fall through to the next case
	switch category := "Teen"; category {
	case "Adult":
		fallthrough
	case "Teen":
		fmt.Println("Age is 13 or more")
	default:
		fmt.Println("Age is less than 13")
	}
}

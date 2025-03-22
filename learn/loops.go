package main

import "fmt"

func loops() {
	fmt.Println("\n# Loops")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// infinite loop
	for {
		fmt.Println("Breaking infinite loop")
		break
	}

	// continue
	fmt.Println("odd numbers")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

}

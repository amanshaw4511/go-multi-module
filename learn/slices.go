package main

import (
	"fmt"
)

// variadic functions
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func slices() {
	fmt.Println("\n# Slices")

	// slice is a dynamic array
	// creating a slice from array
	var a = [5]int{1, 2, 3, 4, 5} // array
	s := a[:]
	fmt.Println(s)

	s1 := a[1:3] // from index 1 to 3 (exclusive)
	fmt.Println(s1)

	// directly creating a slice
	s2 := make([]int, 5)       // length 5, capacity 5
	s3 := make([]int, 5, 10)   // length 5, capacity 10
	s4 := []int{1, 2, 3, 4, 5} // length 5, capacity 5
	fmt.Println(s2, s3, s4)
	fmt.Println("Length of s3:", len(s3))
	fmt.Println("Capacity of s3:", cap(s3))

	// accessing slice elements
	fmt.Println("First element of s3:", s3[0])
	// modifying slice elements
	s3[0] = 10
	fmt.Println("Modified first element of s3:", s3[0])

	// appending to slice
	s3 = append(s3, 6)
	fmt.Println("After appending 6 to s3:", s3)

	// appending multiple elements
	s3 = append(s3, 7, 8, 9)
	fmt.Println("After appending 7, 8, 9 to s3:", s3)

	// appending a slice to a slice
	s3 = append(s3, []int{10, 11, 12}...)
	fmt.Println("After appending [10, 11, 12] to s3:", s3)

	// slicing a slice
	s5 := s3[2:5]
	fmt.Println("s5:", s5)
	all_execpt_first_two := s3[2:]
	fmt.Println("all_execpt_first_two:", all_execpt_first_two)
	last_two := s3[:2]
	fmt.Println("last_two:", last_two)

	// deep copy
	s3_copy := make([]int, len(s3))
	copy(s3_copy, s3)
	fmt.Println("s3_copy:", s3_copy)

	// varaible length arguments( variadic functions)
	fmt.Println("Sum of 1, 2, 3, 4, 5:", sum(1, 2, 3, 4, 5))

	// spread operator
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum of nums:", sum(nums...))

	// range
	for i, v := range nums {
		fmt.Println("Index:", i, "Value:", v)
	}

}

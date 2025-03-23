package main

import "fmt"

func maps() {
	fmt.Printf("\n# Maps\n")

	// create a map
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	// declare and initialize a map
	m1 := map[string]int{"k1": 11, "k2": 12}
	fmt.Println("map1:", m1)

	// get a value
	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	// get the length
	fmt.Println("len:", len(m))
	// delete a value
	delete(m, "k2")
	fmt.Println("map:", m)
	// check if a key exists
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// iterate over a map
	for k, v := range m {
		fmt.Println(k, v)
	}

}

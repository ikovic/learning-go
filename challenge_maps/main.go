package main

import "fmt"

/*
	1. Declare a map called m1 which value is nil. Print out its type and value.
	2. Declare a map called m2. It's keys are of type int and values of type string.  Initialize the map using  a map literal with two key:value pairs.
	3. Add the following key: value to the map: 10: "Abba"
	4. Retrieve the value of an existing key and the value of a non existing key. Think about the results.
*/
func first() {
	var m1 map[int]int
	fmt.Printf("Map m1 %T %v \n", m1, m1)

	m2 := map[int]string{
		1: "First",
		2: "Second",
	}

	m2[10] = "Abba"

	existing := m2[10]
	missing := m2[11]

	fmt.Printf("Map m2 %v\n existing %v missing %v\n", m2, existing, missing)
}

func second() {
	// fix the code
	m1 := map[int]bool{}
	m1[5] = true

	m2 := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}

	_, _, _ = m1, m2, m3
}

/*
	1. The above map declaration has an error. Solve the error!
	2. Delete a key:value pair from the map.
	3. Iterate over the map and print out each key and value.
*/
func third() {
	m := map[int]bool{1: true, 2: false, 3: false}
	delete(m, 2)
	for k, v := range m {
		fmt.Printf("k: %d, v: %t\n", k, v)
	}
}

func main() {
	first()
	second()
	third()
}

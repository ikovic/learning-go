package main

import "fmt"

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string {
	return c.licenceNo
}

func (c car) Name() string {
	return c.brand
}

func first() {
	/*
		1. Create a Go program where car type implements the vehicle interface.
		2. Create a variable of type vehicle that holds a car struct value.
		3. Call the methods (Licence and Name) on the interface value declared at step 2
		4. Run the program without errors.
	*/
	var audi vehicle = car{licenceNo: "ST1234P", brand: "Audi"}
	fmt.Printf("Brand: %v License: %v\n", audi.Name(), audi.License())
}

func second() {
	/*
		1. Declare a variable called empty of type empty interface. Print out its type.
		2. Assign an int value to the variable called empty. Print out its type.
		3. Assign a float64 value to empty. Print out its type.
		4. Assign an int slice value to empty. Print out its type.
		5. Add a new int value to the slice (empty variable).
		6. Print out the slice (empty variable).
	*/

	var empty interface{}
	fmt.Printf("Type of empty: %T\n", empty)
	empty = 6.4
	fmt.Printf("Type of empty: %T\n", empty)
	empty = []int{1, 2, 3}
	fmt.Printf("Type of empty: %T\n", empty)
	empty = append(empty.([]int), 4)
	fmt.Printf("Type of empty and value: %T %v\n", empty, empty)
}

// fix errors
type cube struct {
	edge float64
}

func volume(c cube) float64 {
	return c.edge * c.edge * c.edge
}

func third() {
	var x interface{}
	x = cube{edge: 5}
	v := volume(x.(cube))
	fmt.Printf("Cube Volume: %v\n", v)

}

func main() {
	first()
	second()
	third()
}

package main

import "fmt"

func first() {
	/*
		Consider the following variable declaration x := 10.10
		1. Print out the address of x
		2. Declare a pointer called ptr that stores the address of x.
		3. Print out the type and the value of ptr.
		4. Print the address of the pointer and the value of x though the pointer (use the dereferencing operator).
	*/

	x := 10.10
	fmt.Printf("Address of x: %v\n", &x)

	ptr := &x
	fmt.Printf("Type and value of ptr: %T %v\n", ptr, ptr)

	fmt.Printf("Address of ptr and value of x through dereferencing: %v %v\n", &ptr, *ptr)
}

func second() {
	x, y := 10, 2
	ptrx, ptry := &x, &y

	/* Declare a new variable called z and initialize the variable by dividing x by y through the pointers. */

	z := *ptrx / *ptry
	fmt.Printf("z: %v\n", z)
}

func swap(x *float64, y *float64) {
	*x, *y = *y, *x
}

func third() {
	x, y := 5.5, 8.8
	fmt.Printf("Before swap - x: %v y: %v\n", x, y)
	/* Write a function that swaps the values of x and y. After calling the function x will be 8.8 and y will 5.5 */
	swap(&x, &y)
	fmt.Printf("After swap - x: %v y: %v\n", x, y)

}

func main() {
	first()
	second()
	third()
}

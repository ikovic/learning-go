package main

import "fmt"

func main() {
	const days int = 7
	var i int
	fmt.Println(i)

	const pi float64 = 3.14
	const secondsInHour = 3600

	duration := 234 // in hours
	fmt.Printf("Duration in seconds: %v\n", duration*secondsInHour)

	// x, y := 5, 0
	// fmt.Println(x / y)

	const a = 5
	const b = 0

	// fmt.Println(a / b)
	const n, m int = 4, 5
	const n1, m1 = 6, 7

	const (
		min1 = -500
		min2 = -300
		min3 = 100
	)

	fmt.Println(min1, min2, min3)

	/*
		Rules of constants:
		- you cannot change a constant value
		- you cannot initialize a constant during runtime, only during compilation
		-- exception: len() is a function that is built-in and available during compilation
		- you cannot initialize a constant by pointing to another variable
		- when creating a constant we can make it typed or untyped (behaves different from the regular variables)
		-- untyped constants can be used more flexible (not all type system guarantees are employed)
		--- visible when trying to multiply values of different numerical types, which is otherwise impossible without explicit conversion
	*/

	const (
		c1 = iota
		c2
		c3
	)

	fmt.Println("Iota: ", c1, c2, c3)

	/*
		iota is used as a generator, can be customized by adding operators lika (iota * 2) + 1
		we can use the blank identifier _ to skip elements
	*/

	const (
		x = iota
		_
		y
	)

	fmt.Println("x = 0, y = 2", x == 0 && y == 2)
}

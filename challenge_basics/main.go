package main

import (
	"fmt"
	"math"
)

const version string = "3.1"

func declaration() {
	/*
		#1
		Using the var keyword, declare 3 variables called a, b, c, d of type int, float64, bool and string.
		Using short declaration syntax declare and assign these values to variables x, y and z:
		- 20
		- 15.5
		- "Gopher!"
		Using fmt.Println() print out the values of a, b, c, d, x, y and z.	Try to run the program without error.
	*/
	var (
		a int
		b float64
		c bool
		d string
	)

	x := 20
	y := 15.5
	z := "Gopher!"

	fmt.Println(a, b, c, d, x, y, z)

	/*
		#2
		1. Declare a, b, c, d on a single line for better readability -> multiple declarations
		2. Declare x, y and z on a single line -> multiple short declarations
		3. Remove the statement that prints out the variables. See the error!
		4. Change the program to run without error using the blank identifier (_)
	*/

	x1, y1, z1 := 20, 15.5, "Gopher"
	_, _, _ = x1, y1, z1

	/*
		#3 fix errors
	*/

	var a3 float64 = 7.1

	x3, y3 := true, 3.7

	a3, x3 = 5.5, false

	_, _, _ = a3, x3, y3

	/*
		#4 fix errors
	*/

	name := "Golang"
	fmt.Println(name, version)
}

func constantFixError() {
	const a int = 7
	const b float64 = 5.6
	const c = float64(a) * b

	const x = 8
	const xc int = x

	var noIPv6 = math.Pow(2, 128)
	_ = noIPv6
}

func constant() {
	const (
		daysWeek   = 7
		lightSpeed = 299792458
		pi         = 3.14159
		secPerDay  = 60 * 60 * 24
		daysYear   = 365
	)

	fmt.Printf("Number of seconds in a year is %d\n", secPerDay*daysYear)

	constantFixError()

	const (
		jun = iota + 6
		jul
		aug
	)
}

func types() {
	type duration int
	var hour duration

	fmt.Println(hour)
	fmt.Printf("Type of hour: %T\n", hour)

	hour = 3600

	fmt.Println(hour)
}

func main() {
	declaration()
	constant()
	types()
}

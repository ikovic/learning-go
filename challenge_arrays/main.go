package main

import "fmt"

func first() {
	// Using the var keyword, declare an array called cities with 2 elements of type string. Don't initialize the array.
	var cities = [2]string{}
	fmt.Println(cities)

	// Declare an array called grades of type [3]float64 and initialize only the first 2 elements using an array literal.
	var grades = [3]float64{
		0: 2,
		1: 5,
	}
	fmt.Println(grades)

	// Declare an array called taskDone using the ellipsis operator. The elements are of type bool.
	var taskDone = [...]bool{false, false, true}
	fmt.Println(taskDone)

	// Iterate over the array called cities using the classical for loop syntax and len function and print out element by element.
	for i := 0; i < len(cities); i++ {
		fmt.Printf("Index: %d Value: %s\n", i, cities[i])
	}

	// Iterate over grades using the range keyword and print out element by element.
	for index, value := range cities {
		fmt.Printf("Index: %d Value: %s\n", index, value)
	}
}

func second() {
	// Write a Go program that counts the number of positive even numbers in the array.
	nums := [...]int{30, -1, -6, 90, -6}
	nOfPositiveEvenNumbers := 0

	for _, v := range nums {
		if v > 0 && v%2 == 0 {
			nOfPositiveEvenNumbers++
		}
	}

	fmt.Printf("Number of positive even numbers is %d\n", nOfPositiveEvenNumbers)
}

func third() {
	// fix errors in the code
	myArray := [3]float64{1.2, 5.6}

	myArray[2] = 6

	a := 10
	myArray[0] = float64(a)

	myArray[2] = 10.10

	fmt.Println(myArray)
}

func main() {
	first()
	second()
	third()
}

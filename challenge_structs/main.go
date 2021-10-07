package main

import "fmt"

/*
	1. Create your own struct type called person that stores the following data: name, age and a list with favorite colors.
	2. Declare and initialize two values of type person, one called me and another called you.
	3. Print out the struct values.
*/
func first() {
	type person struct {
		name   string
		age    int
		colors []string
	}

	me := person{name: "Test", age: 15, colors: []string{"red", "green", "blue"}}
	you := person{name: "Andrei", age: 151, colors: []string{"red", "yellow", "purple"}}

	fmt.Printf("Me: %#v You: %#v\n", me, you)
}

/*
	Consider the code from the previous exercise and:
	1. Change the name or the struct value called me to "Andrei"
	2. Take in a new variable the favorites colors of struct value called you. Print out the type and the value of the new variable.
	3. Add a new favorite color to the second struct value called you.
	4. Print out the struct values.
*/
func second() {
	type person struct {
		name   string
		age    int
		colors []string
	}

	me := person{name: "Test", age: 15, colors: []string{"red", "green", "blue"}}
	you := person{name: "Other", age: 151, colors: []string{"red", "yellow", "purple"}}

	me.name = "Andrei"
	yourColors := you.colors

	fmt.Printf("Me: %#v You: %#v\n", me, you)
	fmt.Printf("Your colors %v of type: %T\n", yourColors, yourColors)

	you.colors = append(you.colors, "brown")
	fmt.Printf("You: %#v\n", you)
}

// Iterate and print out the favorite colors of the struct value called me.
func third() {
	type person struct {
		name   string
		age    int
		colors []string
	}

	me := person{name: "Test", age: 15, colors: []string{"red", "green", "blue"}}

	for _, v := range me.colors {
		fmt.Printf("Your favourite color: %#v\n", v)
	}
}

/*
	Change the code from Exercise #1 and:
	1. Create a new struct type called grades with  2 fields: grade and course
	2. Add another field of type grades to person struct type (embedded struct).
	3. Change the initialization of the struct values called me and you to contain information about the grades.
	4. Change the grade and the course of one struct value to "Golang" and 98.
	5. Print out the struct values.
*/
func fourth() {

	type grades struct {
		grade  int
		course string
	}

	type person struct {
		name   string
		age    int
		colors []string
		grades grades
	}

	me := person{name: "Test", age: 15, colors: []string{"red", "green", "blue"}, grades: grades{grade: 4, course: "History"}}
	you := person{name: "Andrei", age: 151, colors: []string{"red", "yellow", "purple"}, grades: grades{grade: 2, course: "History"}}

	you.grades.course, you.grades.grade = "Golang", 98

	fmt.Printf("Me: %#v You: %#v\n", me, you)
}

func main() {
	first()
	second()
	third()
	fourth()
}

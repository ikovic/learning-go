package main

import "fmt"

// Using a for loop and an if statement print out all the numbers between 1 and 50 that divisible by 7.
func first() {
	for i := 1; i <= 50; i++ {
		if i%7 == 0 {
			fmt.Println(i)
		}
	}
}

// Change the code from the previous exercise and use the continue statement to print out all the numbers divisible by 7 between 1 and 50.
func second() {
	for i := 1; i <= 50; i++ {
		if i%7 != 0 {
			continue
		}
		fmt.Println(i)
	}
}

// Change the code from the previous exercise and use the break statement to print out only the first 3 numbers divisible by 7 between 1 and 50.
func third() {
	counter := 0

	for i := 1; i <= 50; i++ {
		if i%7 == 0 {
			fmt.Println(i)
			counter++
		}

		if counter >= 3 {
			break
		}
	}
}

// Using a for loop, an if statement and the logical and operator print out all the numbers between 1 and 500 that divisible both by 7 and 5.
func fourth() {
	for i := 1; i <= 500; i++ {
		if i%7 == 0 && i%5 == 0 {
			fmt.Println(i)
		}
	}
}

/*
	Using a for loop print out all the years from your birthday to the current year.
	Use a variant of for loop where the post statement is moved inside the for block of code.
*/
func fifth() {
	birthYear := 1989

	for birthYear <= 2021 {
		fmt.Println(birthYear)
		birthYear++
	}
}

func sixth() {
	age := 18

	switch {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age < 18:
		fmt.Println("You are minor!")
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")
	}
}

func main() {
	first()
	second()
	third()
	fourth()
	fifth()
	sixth()
}

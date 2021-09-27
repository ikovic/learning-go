package main

import "fmt"

func f() {

}

func main() {
	// int type
	var i1 int = 100
	fmt.Printf("%T\n", i1)

	/*
		byte and rune are aliases used to represent character values
	*/
	var r rune = 'f'
	fmt.Printf("%T\n", r)
	fmt.Println(r)
	fmt.Printf("%x\n", r)

	// arrays
	var numbers = [4]int{1, 2, 3, -4}
	fmt.Printf("%T\n", numbers)

	/* arrays have fixed sizes determined at declaration time, but slices are dynamic instead */
	// slices
	var aSlice = []string{"Paris", "New York", "Peckham"}
	fmt.Printf("%T\n", aSlice) // []string

	// maps (here key type is string, value is float64)
	balances := map[string]float64{
		"USD": 100.52,
		"EUR": 80.2,
	}
	fmt.Printf("%T\n", balances)

	// struct
	type Person struct {
		name string
		age  int
	}

	var you Person
	fmt.Printf("%T\n", you) // main.Person

	// pointer type (stores a memory address of another variable)
	var x int = 2
	ptr := &x
	fmt.Printf("ptr is of type %T with the value of %v\n", ptr, ptr)

	// function type
	fmt.Printf("%T\n", f) // func()

}

package main

import "fmt"

type money float64

func (m money) print() {
	fmt.Printf("%.2f\n", m)
}

func (m money) printStr() string {
	return fmt.Sprintf("%.2f\n", m)
}

func first() {
	/*
		1. Create a new type called money. Its underlying type is float64.
		2. Create a method called print that prints out the money value with exact 2 decimal points.
	*/

	var m money = 6.5

	m.print()
}

func second() {
	/*
		Consider the money type declared at exercise #1.
		Create a new method for the money type called printStr that returns the money value as a string with 2 decimal points.
	*/

	var m money = 6.789

	fmt.Println(m.printStr())
}

type book struct {
	price float64
	title string
}

func (b book) vat() float64 {
	return b.price * 0.09
}

func (b *book) discount() {
	b.price = b.price * 0.9
}

func third() {
	/*
		1. Create a new struct type called book with 2 fields: price and title of type float64 and string.
		2. Create a method for the newly defined type called vat that returns the vat value if vat is 9%.
		3. Create a method called discount that takes a book value as receiver and decreases the price of the book by 10%.
	*/

	var b book = book{title: "Test", price: 25.5}
	fmt.Println(b.vat())
	b.discount()
	fmt.Println(b.price)
}

// method for book type
func (b *book) changePrice(p float64) {
	b.price = p
}

// fix the code
func fourth() {
	// book value
	bestBook := book{title: "The Trial by Kafka", price: 9.9}

	// changing the price
	bestBook.changePrice(11.99)

	fmt.Printf("Book's price:%.2f\n", bestBook.price) // no change
}

func main() {
	first()
	second()
	third()
	fourth()
}

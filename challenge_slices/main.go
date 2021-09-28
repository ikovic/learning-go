package main

import (
	"fmt"
	"os"
	"strconv"
)

func intSlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

/*
	Using a composite literal declare and initialize a slice of type string with 3 elements.
	Iterate over the slice and print each element in the slice and its index.
*/
func first() {
	var strings = []string{"Some", "String", "Sequence"}
	for i, v := range strings {
		fmt.Printf("Index: %d Value: %s\n", i, v)
	}
}

func second() {
	// fix errors
	mySlice := []float64{1.2, 5.6}

	mySlice[1] = 6

	a := 10.0
	mySlice[0] = float64(a)

	mySlice[1] = 10.10

	mySlice = append(mySlice, a)

	fmt.Println(mySlice)
}

func third() {
	var nums = []float64{1.2, 3.23, 4.5}
	nums = append(nums, 10.1)
	nums = append(nums, 4.1, 5.5, 6.6)

	fmt.Println(nums)

	var n = []float64{11.2, 13.2}
	nums = append(nums, n...)

	fmt.Println(nums)
}

func fourth() {
	var args = os.Args[1:]
	var sum, product int = 0, 1

	for _, arg := range args {
		value, _ := strconv.Atoi(arg)
		sum = sum + value
		product = product * value
	}

	fmt.Printf("Args: %v Sum: %d Product: %d\n", args, sum, product)
}

func fifth() {
	// Using a slice expression and a for loop iterate over the slice ignoring the first and the last two elements.
	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	slice := nums[1 : len(nums)-2]
	var sum int
	for _, n := range slice {
		sum += n
	}

	fmt.Printf("Slice: %v Sum: %d \n", slice, sum)
}

func sixth() {
	// Using copy() function create a copy of the slice.
	// Prove that the slices are not connected by modifying one slice and notice that the other slice is not modified.
	friends := []string{"Marry", "John", "Paul", "Diana"}
	newFriends := make([]string, len(friends))

	copy(newFriends, friends)

	friends[0] = "Wrong"
	fmt.Println("Was the new friends list affected?", newFriends[0] == friends[0])
}

func seventh() {
	// Using append() function create a copy of the slice.
	// Prove that the slices are not connected by modifying one slice and notice that the other slice is not modified.
	friends := []string{"Marry", "John", "Paul", "Diana"}
	var newFriends []string
	newFriends = append(newFriends, friends...)

	copy(newFriends, friends)

	friends[0] = "Wrong"
	fmt.Println("Was the new friends list affected?", newFriends[0] == friends[0])
}

func eight() {
	// Using a slice expression and append() function create a new slice called newYears that contains the first 3 and the last 3 elements of the slice.
	// newYears should be []int{2000, 2001, 2002, 2008, 2009, 2010}
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	newYears := []int{}

	newYears = append(years[:3], years[len(years)-3:]...)
	expectedValue := []int{2000, 2001, 2002, 2008, 2009, 2010}
	fmt.Println("Is the output correct? ", intSlicesEqual(newYears, expectedValue))
}

func main() {
	fmt.Println("#1")
	first()
	fmt.Println("#2")
	second()
	fmt.Println("#3")
	third()
	fmt.Println("#4")
	fourth()
	fmt.Println("#5")
	fifth()
	fmt.Println("#6")
	sixth()
	fmt.Println("#7")
	seventh()
	fmt.Println("#8")
	eight()
}

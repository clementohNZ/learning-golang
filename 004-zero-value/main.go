package main

import "fmt"

// default values for different types when they are
// unassigned
// https://tour.golang.org/basics/12
var y string
var z int

func main() {
	fmt.Println("printing the value of y", y)
	fmt.Printf("%T\n", y) // zero value of y is empty string

	y = "Bond, James Bond"

	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T", z)
}
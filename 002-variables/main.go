package main

import "fmt"

func main() {
	// short-declaration operator
	x := 42
	fmt.Println(x)
	x = 49
	fmt.Println(x)

	y := 100  + 24
	fmt.Println(y)

	/*
	var
	---
	differences:
		1. can be declared outside a function
		2. can assign specific type
		3. you can declare with identifier without assignment.
		   A zero value of its type is given by default
		   https://tour.golang.org/basics/12
	*/
	var z = 43
	fmt.Println(z)
}

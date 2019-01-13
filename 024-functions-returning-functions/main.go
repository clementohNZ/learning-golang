package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)

	// returning a function
	f := bar()
	fmt.Println(f) // returns memory reference to function
	fmt.Println(f()) // executes function and returns 451
	fmt.Println(bar()()) // cleaner - not assigning to variable unnecessarily
}

func foo() string {
	return "Hello World"
}

func bar() func() int {
	return func() int {
		return 451
	}
}

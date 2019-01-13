package main

import "fmt"

var x int // global scope variable

/*
Closures
Limiting scopes of variables by enclosing them inside a function or block
You want to do this in order to prevent memory leaks and variable name clashes
 */
func main() {
	fmt.Println(x)

	var y int
	fmt.Println(y)

	{
		z := 2
		fmt.Println(z)
	}
	// fmt.Println(z) - cannot print z here

	// xi - cannot access xi here

	// It references the a inside the incrementor scope
	b := incrementor()
	c := incrementor()
	fmt.Println(b()) // 1
	fmt.Println(b()) // 2
	fmt.Println(b()) // 3
	fmt.Println(b()) // 4
	fmt.Println(c()) // 1 - references different instance of a (in memory)
	fmt.Println(c()) // 2
}

func foo() {
	// y not accessible here since it is in a different scope
	xi := 3
	fmt.Println(xi)
}

func incrementor() func() int {
	a := 0
	return func() int {
		// can access x because x in this function being returned
		// because a is in the parent scope of this function
		a++
		return a
	}
}

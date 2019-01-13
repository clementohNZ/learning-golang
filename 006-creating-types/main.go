package main

import "fmt"

var a int

type hotdog int
var b hotdog

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 43
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// a = b doesn't work (by default)

	// conversion (converting to another type)
	// a.k.a. casting in other languages
	a = int(b)
	fmt.Printf("%T\n", a)
}

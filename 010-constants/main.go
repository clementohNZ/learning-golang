package main

import "fmt"

// const a = 42
// const b = 42.78
// const c = "James Bond"

const (
	a = 42
	b float64 = 42.78
	c string = "James Bond"
	d = iota
	e = iota
	f = iota
)

const (
	g = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// auto increments by 1 starting from
	// the index position of the const variable
	// d is on the 3rd line of the const block and so
	// d is 3, not 0
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)

	// resets to 0 in new const block
	fmt.Println(g)
	fmt.Printf("%T\n", g)
}

package main

import "fmt"

/*
Pointers refer to a place in memory where something is stored

- Useful when dealing with large objects/data so you don't need to copy them and instead just passing around
  the reference
- When you want to change something at a specific address
- Everything in Go is pass by value. If you want to change something at an address, you pass in a copy of that address
  and inside the function you use the address to get the value
 */
func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", a) 	// int
	fmt.Printf("%T\n", &a) 	// *int -> different type (pointer to an int)

	// var b *int = &a
	b := &a
	fmt.Println(b)
	fmt.Println(*b) // * -> get value of whatever is stored at that address
	fmt.Println(&*b) // get address back

	*b = 43
	fmt.Println(a) // a is now equal to 43 because b points to the address of a

	// passing pointers (by value)
	x := 0
	foo(x)
	fmt.Println("original x:", x)

	x1 := 0
	fooManipulateRef(&x1)
	fmt.Println("original x1:", x1)
}

func foo(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}

func fooManipulateRef(y *int) {
	fmt.Println(*y)
	*y = 43
	fmt.Println(*y)
}
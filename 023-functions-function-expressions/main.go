package main

import "fmt"

/*
Functions are first-class citizens in go so they can be assigned to variables
like other types
 */
func main() {
	f := func() {
		fmt.Println("my first func expression")
	}
	f()

	f2 := func(x int) {
		fmt.Printf("The year big brother came out was %d", x)
	}
	f2(1984)
}

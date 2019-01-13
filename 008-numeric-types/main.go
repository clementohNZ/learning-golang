package main

import "fmt"

func main() {
	x := 42
	y := 3.2 // auto float64 inference

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	var a int = 2
	var b float32 = 3.241

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	// you cannot assign a number that's too large for the type
	// e.g. you cannot store 256 in uint8 because it can only store up to 255

	// byte = 8 bits = byte type
}

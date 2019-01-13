package main

import "fmt"

func main() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	// shift the counter 1 place to the left (on the bit numerical system)
	// such that the counters are moved 1 place higher on the bit numerical system
	// e.g.
	// 16s 8s 4s 2s 1
	// 4 = 1 mark in the 4 bits column
	// shifting it to the left means it's now
	// 8 = 1 mark in the 8 bits column
	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	x = 9 // one 8 and one 1
	y = x << 1 // one 16 and one 2
	fmt.Printf("%d\t\t%b\n", y, y) // equals 18

	// kb := 1024
	// mb := 1024 * kb
	// gb := 1024 * mb

	// while this is clever, it's much harder for people
	// to read and wouldn't always be practical in a real-world
	// scenario
	const (
		_ = iota
		kb = 1 << (iota * 10)
		mb = 1 << (iota * 10)
		gb = 1 << (iota * 10)
	)

	fmt.Printf("%d\t\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t\t%b\n", gb, gb)
}

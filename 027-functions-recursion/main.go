package main

import "fmt"

/*
Most things you can do using recursion you can do with loops
which are probably much easier to understand flow-wise
 */
func main() {
	fmt.Println(4 * 3 * 2 * 1)

	n := factorial(4)
	fmt.Println(n)

	n2 := loopFact(4)
	fmt.Println(n2)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopFact(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
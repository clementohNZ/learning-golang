package main

import "fmt"

func main() {
	/*
	For Loop

	Same as javascript and java
		for init; condition; post {}
		ForStmt = "for" [ Condition | ForClause | RangeClause ] Block
	 */
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}

	/*
	Nested Loops
	Try not to do as it increases your Big O Complexity
	score by the power of 1
	 */
	for i := 0; i <= 10; i++ {
		for j := 0; j <= 3; j++ {
			fmt.Printf("The outer loop: %d\t The inner loop: %d\n", i, j)
		}
	}

	/*
	For loop when boolean
	i.e. while loop in other languages
	 */
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}

	/*
	For loop without its own condition
	 */
	y := 1
	for {
		if y > 9 {
			break
		}
		fmt.Println(y)
		y++
	}

	/*
	Continue
	 */
	a := 0
	for {
		a++
		if a > 100 {
			break
		}
		if a%2 != 0 {
			continue
		}
		fmt.Println(a)
	}
}

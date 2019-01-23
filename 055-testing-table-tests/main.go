package main

import "fmt"

func main() {
	fmt.Println("2 + 3 = 5", mySum(2, 3))
	fmt.Println("4 + 6 = 10", mySum(4, 6))
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

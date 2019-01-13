package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 18}
	s := sum(ii...)
	fmt.Printf("Sum of all numbers is %d\n", s)

	se := sumEven(sum, ii...)
	fmt.Printf("Sum of even numbers is %d", se)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

// function as parameter
func sumEven(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

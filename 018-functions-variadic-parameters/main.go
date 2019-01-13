package main

import "fmt"

func main() {
	// sum
	sum1 := sum(2, 3, 4, 5, 6, 7)
	fmt.Printf("The sum is %d", sum1)

	xi := []int{2, 3, 4, 5, 6, 7}
	sum2 := sum(xi...)
	fmt.Println(sum2)

}

/*
Variadic parameters
- Variadic means 0 or more
- Variadic parameter has to be the last parameter
 */
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position, ", i, " we are now adding, ", v, " to the total which is now ", sum)
	}
	fmt.Printf("The total is %d\n", sum)
	return sum
}

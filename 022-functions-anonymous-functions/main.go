package main

import "fmt"

func main() {
	func() {
		fmt.Println("Anonymous function ran")
	}()

	func(x int) {
		fmt.Printf("The meaning of life is %d", x)
	}(42)
}

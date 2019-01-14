package main

import "fmt"

func main() {
	c := make(chan int, 2)

	c <- 42
	c <- 43

	println(<-c)
	println(<-c)

	fmt.Println("---------")

	fmt.Printf("%T\n", c) // chan int

	c2 := make(chan<- int, 2)
	fmt.Printf("%T\n", c2) // chan<- int

	// println(<-c2) -----> doesn't compile

	c3 := make(<-chan int, 2)
	fmt.Printf("%T\n", c3) // <-chan int

	// c3 <- 42 -------> doesn't compile
}

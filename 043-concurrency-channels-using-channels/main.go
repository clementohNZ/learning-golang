package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go foo(c)

	/*
	receive

	If you use a goroutine for bar, there may not be enough time
	for foo to send the value to the channel and bar to receive it
	so your program will exit before it does anything. You need to
	keep this interaction in mind.

	If you want to use a goroutine for bar, you will need to use a
	wait group.
	 */
	// go bar(c)
	bar(c)

	fmt.Println("about to exit")
}

// send
func foo(c chan<- int) {
	c <- 42
}

// receive
func bar(c <-chan int) {
	fmt.Println("The value received from the channel:", <-c)
}

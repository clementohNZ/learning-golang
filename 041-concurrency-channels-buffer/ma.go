package main

import "fmt"

func main() {
	// creating a buffer of 1
	c := make(chan int, 1)

	/*
	Doesn't block execution of code anymore since the
	channel has a buffer of one value. Think of it as a queue.
	The queue isn't over-capacity yet so it can technically keep
	going.

	If I added another thing to the channel, that would block
	code execution
	 */
	c <- 42

	/*
	Since it's not blocking, it can continue to this line
	that pulls off the value from the channel
	 */
	fmt.Println(<-c)
}

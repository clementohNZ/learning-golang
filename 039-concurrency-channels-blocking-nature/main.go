package main

import "fmt"

func main() {
	/*
	fatal error: all goroutines are asleep - deadlock!

	This code doesn't run because the receiver
		fmt.Println(<-c)
	isn't ready to receive when the publisher
		c <- 42
	is ready to hand over the baton (relay racing metaphor). They
	need to happen at the SAME TIME. Channels BLOCK. This is a
	fundamental thing to understand about channels. The blocker
	here is: c <- 42. Since nothing is receiving it, the channel
	will just keep waiting.
	 */
	c := make(chan int)
	c <- 42
	fmt.Println(<-c)

}

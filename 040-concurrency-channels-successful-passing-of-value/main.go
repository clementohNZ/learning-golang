package main

import "fmt"

func main() {
	c := make(chan int)

	/*
	main program reaches this line and hands off the func
	to be executed in another goroutine. That go routine is now
	blocked waiting for something to pull off this channel value
	*/
	go func() {
		c <- 42
	}()

	/*
	main goroutine can continue executing since nothing is blocking it
	so when it reaches this line it pulls off the value from the channel
	and that goroutine is unblocked
	*/
	fmt.Println(<-c)
}

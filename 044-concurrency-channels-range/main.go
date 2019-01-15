package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	/*
	Receive

	Range - range will keep looping until channel is closed
	If we don't close the channel, goroutine will stop because its being
	blocked while waiting for more values
	 */
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
package main

import (
	"fmt"
	"sync"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	// send
	go send(eve, odd) // fan-out

	/*
	receive

	NOTICE:
	See how this receive function is passed to a goroutine
	to execute concurrently whereas in previous examples we didn't
	do this to block the receiving of channel values so the program
	won't exit before it has had enough time to complete.
	 */
	go receive(eve, odd, fanin) // fan-out

	for v := range fanin {
		fmt.Printf("from fan in channel: %d\n", v)
	}

	fmt.Println("about to exit")
}

// send channel
func send(e, o chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

// receive channel
func receive(e <-chan int, o <-chan int, fi chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range e {
			fi <- v // fan-in
		}
		wg.Done()
	}()

	go func() {
		for v := range o {
			fi <- v // fan-in
		}
		wg.Done()
	}()

	wg.Wait()
	close(fi)
}

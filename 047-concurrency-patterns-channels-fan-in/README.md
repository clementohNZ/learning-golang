# Fan-In Pattern
Fan-In is a concurrency design pattern.

In this way you are distributing the workload and collecting all the results when the work is completed.

You can see through the lines `fi <- even` and `fi <- odd` we are taking two channels and moving their values into
a single channel `fi`.

```go
// receive channel
func receive(e, o chan<- int, fi chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for even := range e {
			fi <- even
		}
		wg.Done()
	}()

	go func() {
		for odd := range o {
			fi <- odd
		}
		wg.Done()
	}()

	wg.Wait()
	close(fi)
}
```

This interconnected pattern will take quite some time to get used to.

## Other Examples
[Go Concurrency Patterns: Pipelines and cancellation](https://blog.golang.org/pipelines)

[GO: A BETTER FAN-OUT, FAN-IN EXAMPLE](https://austburn.me/blog/a-better-fan-in-fan-out-example.html)

[Concurrency Patterns: Golang](https://medium.com/@thejasbabu/concurrency-patterns-golang-5c5e1bcd0833)

### Rob Pike's Example
This is the first time we see the double arrow syntax. You are receiving from the channel and immediately sending it
to another channel.

`c <- <-input1`

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// FAN IN
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}
```
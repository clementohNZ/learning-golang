# Channels Select
Go’s select lets you wait on multiple channel operations. Combining goroutines and channels with select is a 
powerful feature of Go.

For our example we’ll select across two channels.

Each channel will receive a value after some amount of time, to simulate e.g. blocking RPC operations executing in 
concurrent goroutines.

We’ll use select to await both of these values simultaneously, printing each one as it arrives.

We receive the values "one" and then "two" as expected.

Note that the total execution time is only ~2 seconds since both the 1 and 2 second Sleeps execute concurrently.

```go
package main

import "time"
import "fmt"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
```


[Go By Example](https://gobyexample.com/select)

[Go Design Patterns](https://www.godesignpatterns.com/)

## Closing Channels
It's OK to leave a Go channel open forever and never close it. When the channel
is no longer used, it will be garbage collected.

Note that it is only necessary to close a channel if the receiver is looking for a close. Closing the
channel is a control signal on the channel indicating that no more data follows.

[Stack Overflow Answer](https://stackoverflow.com/a/8593986/5561773)

I would recommend closing a channel whenever you know there is not going to be any more data flowing through
because it keeps your program purposeful and optimized.
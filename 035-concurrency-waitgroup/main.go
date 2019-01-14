package main

import (
	"fmt"
	"runtime"
	"sync"
)

// primitive way of doing concurrency
var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1) // wait for one thing

	// concurrent but not parallel unless you have more
	// than 1 cpu. foo doesn't print anything because the program
	// continues to run synchronously and exits before foo can do
	// anything.
	go foo()

	bar()

	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	// once its done, decrease amount of things waiting for
	// now foo will print
	wg.Done()
}

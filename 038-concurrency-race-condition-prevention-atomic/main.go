package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/*
Atomic

Package atomic provides low-level atomic memory primitives useful for implementing synchronization algorithms.

https://golang.org/pkg/sync/atomic/
https://gobyexample.com/atomic-counters
 */
func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines (start):", runtime.NumGoroutine())

	var counter int64 // when you see int64, often involves atomic package

	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Goroutines (within loop):", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Goroutines (end):", runtime.NumGoroutine())
	fmt.Println("Count:", counter) // can print without race condition at end
}

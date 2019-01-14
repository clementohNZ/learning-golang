package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Mutex

Used to lock a variable (while one goroutine is using it. Once its no longer used
the variable will be unlocked for other goroutines to use.

Mutex allows you to block r+w, r only, w only, etc.

Different from 036 race condition, this one does't have a race condition and the counter
gets to 100.

https://gobyexample.com/mutexes
 */
func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines (start):", runtime.NumGoroutine())
	counter := 0

	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock() // 1st change
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock() // 2nd change
			wg.Done()
		}()
		fmt.Println("Goroutines (within loop):", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Goroutines (end):", runtime.NumGoroutine())
	fmt.Println("Count:", counter)
}

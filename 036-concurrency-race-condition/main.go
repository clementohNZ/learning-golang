package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines (start):", runtime.NumGoroutine())
	counter := 0

	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < 100; i++ {
		go func() {
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched() // allows something else to run
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines (within loop):", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Goroutines (end):", runtime.NumGoroutine())
	fmt.Println("Count:", counter)
}

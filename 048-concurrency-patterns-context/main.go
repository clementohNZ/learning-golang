package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	//ctx := context.Background()
	//
	//fmt.Println("context:\t", ctx)
	//fmt.Println("context err:\t", ctx.Err())
	//fmt.Printf("context type:\t%T\n", ctx)
	//fmt.Println("----------")
	//
	//ctx, cancel := context.WithCancel(ctx)
	//
	//fmt.Println("context:\t", ctx)
	//fmt.Println("context err:\t", ctx.Err()) // error is nil
	//fmt.Printf("context type:\t%T\n", ctx)
	//fmt.Println("cancel:\t\t", cancel)
	//fmt.Printf("cancel type:\t%T\n", cancel)
	//fmt.Println("----------")
	//
	//cancel()
	//
	//fmt.Println("context:\t", ctx)
	//fmt.Println("context err:\t", ctx.Err()) // error no longer nil
	//fmt.Printf("context type:\t%T\n", ctx)
	//fmt.Println("cancel:\t\t", cancel)
	//fmt.Printf("cancel type:\t%T\n", cancel)
	//fmt.Println("----------")

	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("error check 1:", ctx.Err())
	fmt.Println("num gortins 1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("num gortins 2:", runtime.NumGoroutine())

	fmt.Println("about to cancel context")
	cancel()
	fmt.Println("cancelled context")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 3:", ctx.Err())
	fmt.Println("num gortins 3:", runtime.NumGoroutine())
}

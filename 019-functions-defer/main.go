package main

import "fmt"

/*
-> https://tour.golang.org/flowcontrol/12
A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not
executed until the surrounding function returns.

A really good blog post on defer and real-world usages:
-> https://blog.golang.org/defer-panic-and-recover

It works in a "similar" way to setTimeout in JS where the function call gets pushed to the end
of the current event loop and executes later once the current call stack has been completely
executed.

Another really good blog post with visuals:
-> https://blog.learngoprogramming.com/golang-defer-simplified-77d3b2b817ff

Gotchas with visuals:
-> https://blog.learngoprogramming.com/gotchas-of-defer-in-go-1-8d070894cb01

Go By Example:
-> https://gobyexample.com/defer
 */
func main() {
	// normally executes first, but when "defer" used, bar runs first
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
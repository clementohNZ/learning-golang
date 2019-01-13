package main

import "fmt"

// func (r receiver) identifier(parameters) (returns(s)) { ... }
// everything in Go is PASS BY VALUE
func main() {
	foo()
	bar("James")

	s1 := woo("Moneypenny")
	fmt.Println(s1)

	x, y := mouse("Ian", "Flemming")
	fmt.Println(x)
	fmt.Println(y)
}

func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := false
	return a, b
}
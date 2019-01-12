package main

import "fmt"

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)

	/*
	Format Printing
	https://play.golang.org/p/rqGVNVP5kl

	group #1 - general printing to standard out
		func Print(a ...interface{}) (n int, err error)
		func Printf(format string, a ...interface{}) (n int, err error)
		func Println(a ...interface{}) (n int, err error)t
	group #2 = printing to a string which you can assign to a variable
		func Sprint(a ...interface{}) string
		func Sprintf(format string, a ...interface{}) string
		func Sprintln(a ...interface{}) string
	group #3 - printing to a file or a web server’s response (anything that implements io.Writer interface)
		func Fprint(w io.Writer, a ...interface{}) (n int, err error)
		func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
		func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
	*/

	var a int
	var b = "James Bond"

	s := fmt.Sprint(a, " something more ", b)
	fmt.Println(s)
	s2 := fmt.Sprintf("%v\t%T\t%T\n", "to pass in", a, b)
	fmt.Println(s2)
}

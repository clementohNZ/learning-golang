package main

import "fmt"

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

/*
Receivers
When you want to attach a method to a specific type, you can use receivers
by specifying the type that the function should be attached to at the front of
the method signature.

func (r receiver) identifier(parameters) (returns(s)) { ... }
 */
func (s secretAgent) speak()  {
	fmt.Println("I am", s.first, s.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last: "Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			first: "Theresa",
			last: "May",
		},
		ltk: true,
	}
	sa1.speak()
	sa2.speak()
}

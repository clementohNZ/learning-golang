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

func (s secretAgent) speak()  {
	fmt.Println("I am", s.first, s.last, "- the secret agent speak")
}

func (p person) speak()  {
	fmt.Println("I am", p.first, p.last, "- the person speak")
}

/*
Anything that has the method speak()
is automatically implementing the 'human' interface. This is unlike
other statically typed languages such as Java or C# where you have to
specifically implement the interface.

Interfaces allow polymorphism by expanding what can be accepted by a
specific method.
 */
type human interface {
	speak()
}

func bar(_ human) {
	fmt.Println("I called human")
}

/*
Interfaces
----------
A value can be of more than one type

Links:
https://gobyexample.com/interfaces
https://tour.golang.org/methods/14 - empty interface
https://tour.golang.org/methods/9
https://medium.com/golangspec/interfaces-in-go-part-i-4ae53a97479c
 */
func main() {
	// secret agent 1 is of type: secretAgent AND human
	sa1 := secretAgent{
		person: person{
			first: "James",
			last: "Bond",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last: "Yes",
	}

	sa1.speak()
	bar(sa1)

	p1.speak()
	bar(p1)
}

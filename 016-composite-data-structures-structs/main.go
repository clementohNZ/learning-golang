package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

/*
Create structure of different types
a.k.a. complex data types, aggregate data types
 */
func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Println(p1.first, p1.last)
	fmt.Println(p2.first, p1.last)

	/*
	Embedded Structs
	 */
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}

	// as you can see, the embedded struct got promoted to the same level
	// as the parent struct. We didn't need to use sa1.person.first
	// the only type you have to specify the inner-type is if the outer-type
	// has a property with the same name
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)

	/*
	Anonymous Structs
	Used when you want to eliminate code pollution.
	You could use this when this struct isn't used anywhere else
	 */
	// p2 := person {
	// 	 first: "James",
	// 	 last: "Cameron",
	// 	 age: 48,
	// }
	p3 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Cameron",
		age:   48,
	}

	fmt.Println(p3)
}

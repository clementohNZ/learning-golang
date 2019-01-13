package main

import "fmt"

/*
Aggregate data structure = structure with same type
Composite data structure = different types composed together
 */
func main() {
	/*
	Arrays
	------
	You need to specify the length
	Arrays have a fixed length. The length is part of its type.
	 */
	var a [5]int
	fmt.Println(a)

	a[3] = 42
	fmt.Println(a)

	/*
	Slices
	------
	Group values of the same type together
	Difference is you don't specify the length in the type
	 */
	b := []int{4, 5, 6, 7, 8, 42}
	println(b)

	fmt.Println(len(b))
	fmt.Println(b[0])
	fmt.Println(b[1])
	fmt.Println(b[2])
	fmt.Println(b[3])
	fmt.Println(b[4])
	fmt.Println(b[5])

	for i, v := range b {
		fmt.Println(i, v)
	}

	// slicing a slice
	fmt.Println(b[1])
	fmt.Println(b[0:1])
	fmt.Println(b[1:3])

	// appending to slice
	b = append(b, 77, 88, 99, 144)
	fmt.Println(b)

	c := []int{2, 3, 4}
	b = append(b, c...) // destructuring arrays/slices
	fmt.Println(b)

	// delete from slice
	b = append(b[:2], b[4:]...)
	fmt.Println(b)

	// make
	// https://golang.org/doc/effective_go.html#allocation_make
	// creates a slice of a certain size
	// this is for performance improvements.
	// when you add to a slice, it will copy the old array into a new array
	// and discard the old one. This bypasses that step (up to a certain initialization point)
	d := make([]int, 10, 100)
	fmt.Println(d)

	// multi-dimensional slice
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}

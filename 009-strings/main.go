package main

import "fmt"

func main() {
	// raw (non-interpreted strings) use backticks instead of double quotes
	s := "Hello, playground"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// default encoding is utf-8
	// string is made of a slice of bytes
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		// print utf 8 character
		fmt.Printf("%#U ", s[i])
	}

	fmt.Println("")

	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
}

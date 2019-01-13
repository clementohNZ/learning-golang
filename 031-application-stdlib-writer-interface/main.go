package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello, playground")
	_, err := fmt.Fprintln(os.Stdout, "Hello, playground") // underlying implementation using Writer interface
	if err != nil {
		fmt.Println("Error writing string using Fprint")
	}
	_, err = io.WriteString(os.Stdout, "Hello, playground")
	if err != nil {
		fmt.Println("Error writing string using io.WriteString")
	}
}

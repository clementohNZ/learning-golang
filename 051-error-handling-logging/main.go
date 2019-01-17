package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	defer foo()
	if err != nil {
		//log.Fatal("err happened", err) // code after this doesn't run including deferred foo() call
		fmt.Println("err happened", err)
		log.Println("err happened", err) // same as fmt.Println, but you have a timestamp
		panic("err happened") // deferred function is called before panic causes program to exit
	}
	fmt.Println("hey") // doesn't run because of panic
}

func foo() {
	fmt.Println("when os.Exit() is called, deferred functions don't run")
}

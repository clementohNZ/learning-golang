package main

import "fmt"

func main() {
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}
	if !true {
		fmt.Println("003")
	}
	if !false {
		fmt.Println("004")
	}
	if 2 == 2 {
		fmt.Println("005")
	}
	if 2 != 2 {
		fmt.Println("006")
	}
	if !(2 == 2) {
		fmt.Println("007")
	}

	/**
	Initialization Statement
	 */
	if x := 42; x == 2 {
		fmt.Println("001")
	}
	// fmt.Println(x) - doesn't work because x is only available
	// in the scope of the if block

	y := 42
	if y == 40 {
		fmt.Println("our value was 40")
	} else if y == 41 {
		fmt.Println("our value was 41")
	} else {
		fmt.Println("our value was not 40 or 41")
	}

	/*
	Switch Statements

	Unlike JavaScript you don't have to put break after each case because
	fallthrough is disabled by default
	 */
	 fmt.Println("-------------------")
	const testString = "John"
	switch testString {
	case "":
		fmt.Println("this should not print")
	case "Jake":
		fmt.Println("this should not print 2")
	case "Melissa":
		fmt.Println("prints")
		fallthrough
	case "John":
		fmt.Println("also true, does it print?")
	default:
		fmt.Println("this is the default")
	}
}

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	n, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println("Error printing text")
	}
	fmt.Println(n)

	// ------- EXAMPLE 2 ---------------

	var answer1, answer2, answer3 string

	fmt.Print("Name: ")
	_, err = fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Food: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Sport: ")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		panic(err)
	}

	fmt.Println(answer1, answer2, answer3)

	// ------- EXAMPLE 3 ---------------

	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("James Bond")

	_, err = io.Copy(f, r)
	if err != nil {
		fmt.Println("Error printing text to file")
	}

	// ------- EXAMPLE 4 ---------------

	f, err = os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}

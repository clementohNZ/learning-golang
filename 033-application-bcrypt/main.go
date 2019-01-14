package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

/*
Use: go get golang.org/x/crypto/bcrypt
to download the bcrypt package
 */
func main() {
	s:= `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)
	fmt.Println(string(bs))

	loginPword := `password1234`
	hashedPword, err := bcrypt.GenerateFromPassword([]byte(loginPword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	err = bcrypt.CompareHashAndPassword(hashedPword, []byte(loginPword))
	if err != nil {
		fmt.Println("You CAN'T login!")
		return
	} else {
		fmt.Println("You're logged in!")
	}
}

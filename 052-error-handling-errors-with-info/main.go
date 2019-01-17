package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		//return 0, errors.New("norgate math: square root of a negative number")
		// fmt.Errorf is just a convenience method for calling errors.New
		return 0, fmt.Errorf("norgate math: square root of a negative number")
	}
	return 42, nil
}

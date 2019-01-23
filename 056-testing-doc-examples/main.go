package main

import (
	"fmt"
	"github.com/clementohNZ/learn-go/056-testing-doc-examples/custommath"
)

func main() {
	fmt.Println(custommath.Sum(2, 3))
	fmt.Println(custommath.Sum(2, 3, 4234, 234, 232, 342, 34))
}

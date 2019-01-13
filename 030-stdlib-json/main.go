package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

type person2 struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}


func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	// marshal/encode json
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	// unmarshal/decode json
	str2Decode := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs2 := []byte(str2Decode)

	var decoded []person2
	err = json.Unmarshal(bs2, &decoded)
	if err != nil {
		fmt.Println(err)
	}
	for i, v := range decoded {
		fmt.Printf("Person Number %d\n", i)
		fmt.Printf("First name is: %v\n", v.First)
	}
}

package main

import "fmt"

func main() {
	/*
	Key value stores for quick lookup
	---------------------------------
	- Maps are unordered
	 */
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"]) // returns zero value when doesn't exist

	// more reliable way to check if key actually exists
	// this is called the "comma okay" idiom
	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("This is the if print", v)
	}

	// adding to map
	m["todd"] = 33
	fmt.Println(m)

	// looping over map
	for k, v := range m {
		fmt.Printf("Key is: %v\t Value is: %v\n", k, v)
	}

	// delete from map
	delete(m, "James")
	fmt.Println(m)

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("value:", v)
		delete(m, "Miss Moneypenny")
	}
	fmt.Println(m)
}

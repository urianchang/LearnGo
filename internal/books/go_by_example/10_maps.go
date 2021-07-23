package main

import "fmt"

func main() {
	// Maps are Go's built-in associative data type (e.g. hashes, dicts)
	m := make(map[string]int)	// map[key-type]val-type
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	fmt.Println("v1:", m["k1"])
	fmt.Println("len:", len(m))

	// delete removes key/value pairs from a map
	delete(m, "k2")
	fmt.Println("deletion:", m)

	// optional second return value when getting a value from a map
	// indicates if the key was present in the map
	_, present := m["k3"]
	fmt.Println("present:", present)

	// declare and init in single line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n:", n)
}
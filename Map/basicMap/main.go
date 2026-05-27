package main

import "fmt"

func main() {

	// Decleare Map
	var mp map[string]int
	fmt.Println(mp)

	// Initialized map
	mp = map[string]int{
		"id":  1,
		"age": 24,
	}
	fmt.Println(mp) // Print last to first

	// Accessing map
	fmt.Println(mp["age"])

	// Modifing map
	mp["age"] = 23
	fmt.Println(mp["age"])

	// Adding new key:value
	mp["balance"] = 2000
	fmt.Println(mp["balance"])

	// Iterative over a map
	for k, v := range mp {
		fmt.Printf("Key: %s -> Value: %d\n", k, v)
	}

	// Delete key:value
	delete(mp, "age")
	fmt.Println(mp)
}

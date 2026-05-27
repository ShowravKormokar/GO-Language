package main

import "fmt"

func main() {
	var p *int
	p = new(int) // Pick an address from memory(random)
	*p = 77
	fmt.Println(p, *p)
}

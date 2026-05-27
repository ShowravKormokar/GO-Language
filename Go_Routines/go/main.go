package main

import (
	"fmt"
	"time"
)

func printNums() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond) // set time interval for check concurrent exec.
		fmt.Println(i)
	}
}

func printChars() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(300 * time.Millisecond) // set time interval for check concurrent exec.
		fmt.Printf("%c\n", i)
	}
}

func main() {
	// This 2 function run one by one (sequentially) under main func go-routine
	// printNums()
	// printChars()

	// Using seperate go routine for this 2 function execution (concurrently execution)
	go printNums()
	go printChars()
	// Print concurrently like Parallel exceution

	time.Sleep(2000 * time.Millisecond)
	fmt.Println("All Execution End")
}

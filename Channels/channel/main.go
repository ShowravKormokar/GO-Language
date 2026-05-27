package main

import "fmt"

func multiply(a int, b int, ch chan int) {
	ch <- a * b // send the result of multiplication to the channel
}

func main() {
	ch := make(chan int) // create a channel of type int

	go multiply(2, 3, ch) // start a goroutine to perform multiplication and send the result to the channel

	res := <-ch // receive the result from the channel

	fmt.Println("Res:", res)
}

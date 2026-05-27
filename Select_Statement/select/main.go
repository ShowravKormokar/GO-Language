package main

import (
	"fmt"
	"time"
)

func server1(ch1 chan string) {
	time.Sleep(2 * time.Second)
	ch1 <- "From Server 1"
}

func server2(ch2 chan string) {
	time.Sleep(4 * time.Second)
	ch2 <- "From Server 2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)

	go server1(output1)
	go server2(output2)

	select { // select statement is used to wait on multiple channel operations. It blocks until one of the channels is ready for communication, and then it executes the corresponding case.
	case s1 := <-output1: // if output1 channel is ready, it will execute this case and print the value received from output1 channel
		fmt.Println("S1:", s1)

	case s2 := <-output2: // if output2 channel is ready, it will execute this case and print the value received from output2 channel
		fmt.Println("S2:", s2)

		// default:// if neither of the channels is ready, it will execute the default case
		// 	fmt.Println("No server responded")
	}
}

/* Here server1 will send data to output1 channel after 2 seconds, and server2 will send data to output2 channel after 4 seconds. The select statement will wait for either of the channels to receive data and print the corresponding message. In this case, it will print "S1: From Server 1" after 2 seconds, and then it will exit without waiting for server2 to send data.
 */

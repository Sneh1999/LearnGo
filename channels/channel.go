package main

import "fmt"

func main() {

	// to make a channel we use make
	messages := make(chan string)
	// the <- is used  to send values to the channel
	go func() { messages <- "ping" }()

	// this is to get the the values from the channel
	msg := <-messages
	fmt.Println(msg)
}

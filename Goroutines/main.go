package main

import (
	"fmt"
	"time"
)

func printMessage(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(800 * time.Millisecond)
	}
	// Sending data to the channel
	channel <- "Done!"
}

func main() { // main goroutine
	// Creating a channel
	channel := make(chan string)

	go printMessage("Go is great!", channel)
	go printMessage("Frontend Masters", channel)
	response := <-channel
	fmt.Println(response)
	// go printMessage("We miss classes!", channel)
}

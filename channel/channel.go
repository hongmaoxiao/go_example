// Use of the select statement with channels, for timeouts, etc.
package main

import (
	"fmt"
	"time"
)

// Function that is "chatty"
//Takes a single parameter a channel to send messages down
func chatter(chatChannel chan<- string) {
	// loop ten times and die
	time.Sleep(3 * time.Second) // sleep for 5 seconds
	chatChannel <- fmt.Sprintf("This is pass number %d of chatter", 1)
}

// out main function
func main() {
	// Create the channel, it will be taking only strings, no need for a buffer on this project
	chatChannel := make(chan string)
	// Clean up our channel when we are done
	defer close(chatChannel)

	// start a go routine with chatter (separate, no blocking)
	go chatter(chatChannel)

	// select statement will block this thread until one of the two conditions below is met
	// because we have a default, we will hit default any time the chatter isn't chatting
	select {
	// anytime the chatter chats, we'll catch it and output it
	case spam := <-chatChannel:
		fmt.Println(spam)
	// if the chatter takes more than 3 seconds to chat, stop waiting
	case <-time.After(3 * time.Second):
		fmt.Println("Ain't no time for that!")
	}
}

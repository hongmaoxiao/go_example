// Use of the select statement with channels (no timeouts)
package main

import (
	"fmt"
	"time"
)

// Function that is "chatty"
// Takes a single parameter a channel to send messages down
func chatter(chatChannel chan<- string) {
	// Clean up our channel when we are done.
	// The channel writer should always be the one to close a channel.
	defer close(chatChannel)

	// loop five times and die
	for i := 1; i <= 5; i++ {
		time.Sleep(2 * time.Second) // sleep for 2 seconds
		chatChannel <- fmt.Sprintf("This is pass number %d of chatter", i)
		fmt.Println("iiii: ", i)
	}
}

// Our main function
func main() {
	// Create the channel
	chatChannel := make(chan string, 1)

	// start a go routine with chatter (separate, non blocking)
	go chatter(chatChannel)

	// This for loop keeps things going while the chatter is sleeping
	for {
		// select statement will block this thread until one of the two conditions below is met
		// because we have a default, we will hit default any time the chatter isn't chatting
		select {
		// anytime the chatter chats, we'll catch it and output it
		case spam, ok := <-chatChannel:
			fmt.Println("chatChannel: ", spam, ok)
			// Print the string from the channel, unless the channel is closed
			// and we're out of data, in which case exit.
			if ok {
				fmt.Println(spam)
			} else {
				fmt.Println("Channel closed, exiting!")
				return
			}
		default:
			// print a line, then sleep for 1 second.
			fmt.Println("Nothing happened this second.")
			time.Sleep(1 * time.Second)
		}
	}
}

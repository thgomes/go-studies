package main

import "fmt"

// thread 1
func main() {
	channel := make(chan string) // canal vazio

	// thread 2
	go func() {
		channel <- "hello world" // enche canal
	}()

	msg := <-channel // esvazia canal
	fmt.Println(msg)
}

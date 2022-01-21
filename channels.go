package main

import ("fmt")

func main() {
	//default is unbuffered
	messages := make(chan string)

	go func() {
		messages <- "ping"
		messages <- "hello"
	}()

	msg := <-messages
	msg2 := <-messages

	fmt.Println(msg)
	fmt.Println(msg2)
}
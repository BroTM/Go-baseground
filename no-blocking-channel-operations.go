package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	//messages <- "hi"
	select {
	case msg := <-messages:
		fmt.Println("message received", msg)
	default :
		fmt.Println("no message received")
	}

	msg := "hi"
	
	select {
	case messages <- msg:
		fmt.Println("messae sent", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
		
	}

}
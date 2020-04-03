package main

import (
	"fmt"
	"time"
)

func send(ch chan<- bool) {
	time.Sleep(120 * time.Millisecond)
	ch <- true
	close(ch)
	fmt.Println("Sent and closed")
}

func main() {
	ch := make(chan bool)
	timeout := time.After(600 * time.Millisecond)

	go send(ch)
	for {
		select {
		case <- ch:
			fmt.Println("Got message")
		case <- timeout:
			fmt.Println("Time out")
			return
		default:
			fmt.Println("*yawn*")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

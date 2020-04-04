package main

import (
	"fmt"
	"time"
)

func main() {
	lock := make(chan bool, 1)
	for i := 1; i < 7; i++ {
		go worker(i, lock)
	}
	time.Sleep(10 * time.Second)
}

func worker(id int, lock chan bool) {
	fmt.Println(fmt.Sprintf("%d wants the lock", id))
	lock <- true
	fmt.Println(fmt.Sprintf("%d has the lock", id))
	time.Sleep(500 * time.Millisecond)
	fmt.Println(fmt.Sprintf("%d is releasing the lock", id))
	<- lock
}
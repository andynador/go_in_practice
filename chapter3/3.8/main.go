package main

import "time"

func main() {
	sleep := time.After(5 * time.Second)

	<- sleep
}
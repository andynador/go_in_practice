package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.DialTimeout("udp", "localhost:1902", 30 * time.Second)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		os.Exit(1)
	}
	defer conn.Close()
	logger := log.New(conn, "example ", log.Ldate | log.Lshortfile)
	logger.Println("This is a regular message")
	logger.Panicln("This is panic")
}

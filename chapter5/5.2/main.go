package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logfile, _ := os.Create("log.txt")
	defer func() {
		if err := logfile.Close(); err != nil {
			fmt.Println(fmt.Sprintf("Error closing file: %s", err.Error()))
			os.Exit(1)
		}
	}()
	logger := log.New(logfile, "example", log.LstdFlags|log.Lshortfile)

	logger.Println("This is a regular message.")
	logger.Fatalln("This is a fatal error")
	logger.Println("This is the end of function")
}

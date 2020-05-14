package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
)

func main() {
	priority := syslog.LOG_LOCAL3 | syslog.LOG_NOTICE
	flags := log.Ldate | log.Lshortfile
	logger, err := syslog.NewLogger(priority, flags)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		os.Exit(1)
	}
	logger.Println("This is a test log message")
}

package main

import (
	"fmt"
	"log/syslog"
	"os"
)

func main() {
	logger, err := syslog.New(syslog.LOG_LOCAL3, "narwhal")
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("Debug message")
	logger.Notice("Notice message")
	logger.Warning("Warning message")
	logger.Alert("Alert message")
}

package main

import (
	"fmt"
	"gopkg.in/gcfg.v1"
	"os"
)

func main() {
	config := struct {
		Section struct{
			Enabled bool
			Path string
		}
	}{}
	err := gcfg.ReadFileInto(&config, "conf.ini")
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to parse ini file: %s", err))
		os.Exit(1)
	}
	fmt.Println(config.Section.Enabled)
	fmt.Println(config.Section.Path)
}

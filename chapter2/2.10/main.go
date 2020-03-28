package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type configuration struct {
	Enabled bool
	Path string
}

func main() {
	file, err := os.Open("conf.yaml")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	decoder := yaml.NewDecoder(file)
	conf := configuration{}
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("%+v", conf))
}

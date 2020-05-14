package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	file, err := OpenCSV("data.csv")
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	defer file.Close()
}

func OpenCSV(filename string) (file *os.File, err error) {
	defer func() {
		if r := recover(); r != nil {
			file.Close()
			err = r.(error)
		}
	}()
	file, err = os.Open(filename)
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to open file: %s", err.Error()))
		return
	}
	RemoveEmptyList(file)

	return
}

func RemoveEmptyList(f *os.File) {
	panic(errors.New("failed parse"))
}
package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(fmt.Sprintf("Trapped panic: %s (%T)", err, err))
		}
	}()
	yikes()
}

func yikes() {
	panic(errors.New("something bad happened"))
}

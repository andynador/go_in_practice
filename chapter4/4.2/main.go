package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func Contact(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("no strings supplied")
	}
	return strings.Join(parts, " "), nil
}

func main() {
	args := os.Args[1:]
	if result, err := Contact(args...); err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err))
		os.Exit(1)
	} else {
		fmt.Println(fmt.Sprintf("Concatenated string: %s", result))
	}
}
package main

import (
	"fmt"
	"log"
	"strings"
)

func Pad(s string, max uint) string {
	log.Println(fmt.Sprintf("Testing Len: %d, Str: %str", max, s))
	ln := uint(len(s))
	if ln > max {
		return s[:max-1]
	}
	s += strings.Repeat(" ", int(max - ln))
	return s
}


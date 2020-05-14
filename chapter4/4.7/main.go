package main

import "fmt"

type ParseError struct {
	Message string
	Line, Char int
}

func (p *ParseError) Error() string {
	return fmt.Sprintf("%s on Line %d, Char %d", p.Message, p.Line, p.Char)
}



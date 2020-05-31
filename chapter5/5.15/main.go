package main

import (
	"io"
)

type MyWriter struct {
}

func (m *MyWriter) Write(bytes []byte) error {
	return nil
}

func main() {
	_ = map[string]interface{} {
		"w": &MyWriter{},
	}

}

func doSomething(m map[string]interface{}) {
	_ = m["w"].(io.Writer)
}
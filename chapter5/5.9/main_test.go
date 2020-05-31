package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	if v := Hello(); v != "Hello" {
		t.Error(fmt.Sprintf("expected 'Hello', but got '%s'", v))
	}
}
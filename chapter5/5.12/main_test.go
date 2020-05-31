package main

import (
	"fmt"
	"testing"
)

type MockMessage struct {
	email, subject string
	body []byte
}

func (m *MockMessage) Send(email, subject string, body []byte) error {
	m.email = email
	m.subject = subject
	m.body = body

	return nil
}

func TestAlert(t *testing.T) {
	msgr := MockMessage{}
	body := []byte("Critical error")
	Alert(&msgr, body)

	if msgr.subject != "Critical error" {
		t.Error(fmt.Sprintf("expected 'Critical error', but got '%s'", msgr.subject))
	}
}
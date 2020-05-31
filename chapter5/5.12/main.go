package main

type Messager interface {
	Send(email, subject string, body []byte) error
}

type Message struct {}

func (m *Message) Send(email, subject string, body []byte) error {
	return nil
}

func Alert(m Messager, problem []byte) error {
	return m.Send("noc@example.com", "Critical error", problem)
}
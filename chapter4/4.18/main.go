package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
)

func main() {
	listen()
}

func listen() {
	listener, err := net.Listen("tcp", ":1026")
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to open port 1026: %s", err.Error()))
		os.Exit(1)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(fmt.Sprintf("Error acception connection: %s", err.Error()))
			os.Exit(1)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	reader := bufio.NewReader(conn)
	data, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to read from socket: %s", err.Error()))
		conn.Close()
	}
	response(data, conn)
}

func response(data []byte, conn net.Conn) {
	panic(errors.New("failure in response"))
}
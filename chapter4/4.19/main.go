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
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(fmt.Sprintf("Fatal error: %s", err))
		}
		conn.Close()
	}()
	reader := bufio.NewReader(conn)
	data, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to read from socket: %s", err.Error()))
	}
	response(data, conn)
}

func response(data []byte, conn net.Conn) {
	conn.Write(data)
	panic(errors.New("pretend I'm a real error"))
}
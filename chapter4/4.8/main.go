package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var ErrTimeout = errors.New("the request timed out")
var ErrRejected = errors.New("the request was rejected")

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func SendRequest(req string) (string, error) {
	switch random.Int() % 3 {
	case 0:
		return "success", nil
	case 1:
		return "", ErrRejected
	default:
		return "", ErrTimeout
	}
}

func main () {
	response, err := SendRequest("hello")
	for err == ErrTimeout {
		fmt.Println("Time out. Retrying")
		response, err = SendRequest("hello")
	}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}

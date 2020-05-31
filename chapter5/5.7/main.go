package main

import "runtime/debug"

func main() {
	foo()
}

func foo() {
	i := 1
	bar(i)
}

func bar(i int) {
	debug.PrintStack()
}
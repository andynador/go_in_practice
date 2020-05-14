package main

import (
	"fmt"
	"sync"
)

func main () {
	var (
		wg sync.WaitGroup
	)
	ch := make(chan int)
	done := make(chan bool)
	count := 60

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(i int) {
			ch <- i
			wg.Done()
		}(i)
	}
	go func() {
		for {
			select {
			case i := <- ch:
				fmt.Println(fmt.Sprintf("Got number %d", i))
				count--
				if count == 0 {
					done <- true
					return
				}
			}
		}
	}()
	go func() {
		wg.Add(1)
		for {
			select {
			case <-done:
				fmt.Println("Goodbye")
				wg.Done()
				return
			}
		}
	}()
	wg.Wait()
	close(ch)
	close(done)
}

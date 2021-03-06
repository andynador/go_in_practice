package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)
func handleError(message string, err error) {
	if err != nil {
		fmt.Println(fmt.Sprintf(message, err.Error()))
		os.Exit(1)
	}
}

func close(errorMessage string, closer io.Closer) {
	err := closer.Close()
	handleError(errorMessage, err)
}

type words struct {
	sync.Mutex
	found map[string]int
}


func newWords() *words {
	return &words{found: make(map[string]int)}
}

func (w *words) add(word string, n int)  {
	w.Lock()
	defer w.Unlock()
	count, ok := w.found[word]
	if !ok {
		w.found[word] = n
		return
	}

	w.found[word] = count + n
}

func tailyWords(filename string, dict *words) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer close("Error during close file: %s", file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict.add(word, 1)
	}

	return scanner.Err()
}

func main()  {
	var wg sync.WaitGroup
	w := newWords()
	for _, f := range os.Args[1:] {
		wg.Add(1)
		go func(file string) {
			if err := tailyWords(file, w); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(f)
	}
	wg.Wait()
	fmt.Println("Words that appear more than once:")
	for word, count := range w.found {
		if count > 1 {
			fmt.Println(fmt.Sprintf("%s: %d", word, count))
		}
	}
}



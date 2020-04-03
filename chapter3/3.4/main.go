package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

func handleError(message string, err error) {
	if err != nil {
		fmt.Println(fmt.Sprintf(message, err.Error()))
		os.Exit(1)
	}
}

func main() {
	var (
		wg sync.WaitGroup
		i  int
		file string
	)
	for i, file = range os.Args[1:] {
		wg.Add(1)
		go func(filename string) {
			err := compress(filename)
			handleError("Error during compress: %s", err)
			wg.Done()
		}(file)
	}
	wg.Wait()
	fmt.Println(fmt.Sprintf("Compressed %d files", i + 1))
}

func close(errorMessage string, closer io.Closer) {
	err := closer.Close()
	handleError(errorMessage, err)
}

func compress(filename string) error {
	var err error
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer close("Error during close in: %s", in)
	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer close("Error during close out: %s", out)

	gzout := gzip.NewWriter(out)
	defer close("Error during close gzout: %s", gzout)

	_, err = io.Copy(gzout, in)
	if err != nil {
		return err
	}

	return nil
}

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]

	f, err := os.Open(fileName)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	_, copyErr := io.Copy(os.Stdout, f)

	if copyErr != nil {
		fmt.Printf("Error copying file contents: %v\n", copyErr)
		return
	}
}

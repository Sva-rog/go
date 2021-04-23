package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing file name!")
		return
	}
	data, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error, can't read file:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, data)
}

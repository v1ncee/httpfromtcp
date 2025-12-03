package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}

	defer f.Close() // Closes when main closes

	buf := make([]byte, 8)

	for {
		n, err := f.Read(buf)

		if n > 0 {
			fmt.Println("read:", string(buf[:n]))
		}

		if err == io.EOF {
			return
		}

		if err != nil {
			fmt.Println("Error reading file", err)
			return
		}
	}
}

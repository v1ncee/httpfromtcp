package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)

	go func() {
		defer close(lines) // close the channel when done
		defer f.Close()

		buf := make([]byte, 8)
		currentLine := ""

		for {
			n, err := f.Read(buf)

			if n > 0 {
				chunk := string(buf[:n])
				parts := strings.Split(chunk, "\n")

				for _, part := range parts[:len(parts)-1] {
					fmt.Printf("read: %s\n", currentLine+part)
					currentLine = ""
				}

				currentLine += parts[len(parts)-1]
			}

			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Println("Error reading file", err)
				return
			}
		}

		if currentLine != "" {
			fmt.Printf("read: %s\n", currentLine)
		}
	}()

	return lines
}

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}

	for line := range getLinesChannel(f) {
		fmt.Printf("read: %s\n", line)
	}
}

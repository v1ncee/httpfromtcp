package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)

	go func() {
		defer close(lines)
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
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		fmt.Println("Error starting listener", err)
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection", err)
			continue
		}

		fmt.Println("New connection accepted from", conn.RemoteAddr())

		for line := range getLinesChannel(conn) {
			fmt.Printf("read: %s\n", line)
		}
	}
}

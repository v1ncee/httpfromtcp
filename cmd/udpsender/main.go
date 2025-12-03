package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		fmt.Println("Error resolving address", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error dialing UDP", err)
		return
	}

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("UDP sender ready.")

	for {
		fmt.Print("> ")

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading line", err)
			continue
		}

		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("Error writing line", err)
			continue
		}

	}
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

// simple telnet server for testing
func main() {

	fmt.Println("Launching server...")

	ln, _ := net.Listen("tcp", ":8081")

	conn, _ := ln.Accept()

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println("error:", err)
			}

			break
		}
		fmt.Print("Message Received:", string(message))
		conn.Write([]byte(message + "\n"))
	}
}

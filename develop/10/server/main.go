package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Starting server...")

	ln, _ := net.Listen("tcp", ":4125")

	conn, _ := ln.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if message != "" {
			fmt.Print("Message Received:", string(message))
			newmessage := strings.ToUpper(message)
			conn.Write([]byte(newmessage + "\n"))
		}
		continue

	}
}

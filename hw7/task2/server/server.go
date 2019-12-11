package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func inputProcessing(msg string) string {
	num, err := strconv.ParseInt(msg, 10, 64)
	if err != nil {
		return strings.ToUpper(msg)
	}
	return strconv.FormatInt(num*2, 10)
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":8081")
	conn, _ := ln.Accept()
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", string(message))
		conn.Write([]byte(inputProcessing(strings.TrimSuffix(message, "\n")) + "\n"))
	}
}

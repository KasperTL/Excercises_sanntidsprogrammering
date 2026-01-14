package main

import (
	"fmt"
	"log"
	"net"
)

func listening_tcp(conn net.Conn) {
	buf := make([]byte, 1024)
	var message []byte

	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("Read error: %v", err)
			return
		}

		message = append(message, buf[:n]...)

		for {
			idx := -1
			for i, b := range message {
				if b == 0 {
					idx = i
					break
				}
			}
			if idx == -1 {
				break
			}

			fmt.Printf("Got message: %s\n", string(message[:idx]))
			message = message[idx+1:]
		}
	}
}

func writing_tcp(conn net.Conn) {
	message := []byte("Connect to: 10.100.23.11:20017\x00")
	_, err := conn.Write(message)
	if err != nil {
		fmt.Println("Write error:", err)
		return
	}
}

func main() {

	serverIP := "10.100.23.11"
	portNumber := "33546"

	raddr, err := net.ResolveTCPAddr("tcp", serverIP+":"+portNumber)
	if err != nil {
		fmt.Println("Resolve error:", err)
		return
	}
	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		fmt.Println("Dial error:", err)
		return
	}
	fmt.Println("Connected to server:", conn.RemoteAddr())

	defer conn.Close()

	go listening_tcp(conn)
	go writing_tcp(conn)

	select {}
}

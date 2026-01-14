package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func listening_tcp(conn net.Conn) {

	welcome_buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(welcome_buffer)
		if err != nil {
			log.Printf("Read error: %v", err)
		}
		fmt.Printf("Got message from %s", string(welcome_buffer[:n]))
	}
}

func writing_tcp(conn net.Conn) {
	for {
		message := []byte("Test group 17")
		_, err := conn.Write(message)
		if err != nil {
			fmt.Println("Write error:", err)
			return
		}
		time.Sleep(time.Second)
	}
}

func main() {

	serverIP := "10.100.23.11"
	portNumber := "34933"

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

	defer conn.Close()

	go listening_tcp(conn)
	go writing_tcp(conn)

	select {}
}

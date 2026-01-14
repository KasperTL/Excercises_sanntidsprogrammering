package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	serverIP := "192.168.1.42" // replace with discovered IP from UDP
	port := 34933

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIP, port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Connected to TCP server (fixed-size messages)")

	// Receive fixed-size welcome message
	welcome := make([]byte, 1024)
	_, err = conn.Read(welcome)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Welcome message:", string(welcome))

	// Send a test message
	message := []byte("Hello TCP server!" + string(make([]byte, 1024-len("Hello TCP server!")))) // pad to 1024 bytes
	_, err = conn.Write(message)
	if err != nil {
		log.Fatal(err)
	}

	// Receive echo
	echo := make([]byte, 1024)
	_, err = conn.Read(echo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Echo:", string(echo))
}

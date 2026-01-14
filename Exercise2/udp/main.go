package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	//Setting up the UDP adress
	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:30000")
	if err != nil {
		log.Fatal("Couldnt resolve address:", err)
	}

	// Start listening
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal("Listen failed:", err)
	}
	defer conn.Close()

	// Buffer for incoming data
	buffer := make([]byte, 1024)
	for {
		// Read client message
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Read error: %v", err)
			continue
		}
		fmt.Printf("Got message from %s: %s\n", clientAddr, string(buffer))

		// Echo back
		_, err = conn.WriteToUDP(buffer[:n], clientAddr)
		if err != nil {
			log.Printf("Write error: %v", err)
		}
	}
}

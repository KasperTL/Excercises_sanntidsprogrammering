package main

import (
	"fmt"
	"net"
)

const delimitedMsgPort = "33546"

func writeToServerTCP() {
	// Connect to the server on port 33546
	conn, err := net.Dial("tcp", "10.100.23.11:"+delimitedMsgPort)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Send data to the server ending with '\000' for delimited messages
	message := "Hello, server! This is group 17.\000"
	conn.Write([]byte(message))
}

func readFromServerTCP() {
	// Connect to the server on port 33546
	conn, err := net.Dial("tcp", "10.100.23.11:"+delimitedMsgPort)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Read the server's response
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	// Print the server's response
	fmt.Printf("Server response: %s\n", buffer[:n])

	n, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	// Print the server's response
	fmt.Printf("Server response: %s\n", buffer[:n])
}

func main() {
	// Run the writeToServerTCP and readFromServerTCP concurrently
	go writeToServerTCP()
	go readFromServerTCP()

	// Sleep to keep the program running
	select {}
}

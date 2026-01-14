package main

import (
	"fmt"
	"log"
	"net"
)

func listenServerUDP() {
	//Setting up the UDP adress
	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:20017")
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

func writeToServerUDP() {
	raddr, err := net.ResolveUDPAddr("udp", "10.100.23.11:20017")
	if err != nil {
		fmt.Println("Resolve error:", err)
		return
	}
	//resolve a string network address into a *netUDPAddr structure
	//requireed by net.DialAddr and net.ListenUDP to establish UDP connection

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		fmt.Println("Dial error:", err)
		return
	}
	defer conn.Close()

	message := []byte("Test group 17")
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Write error:", err)
		return
	}

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Printf("Recived from %s: %s\n", raddr, string(buffer[:n]))
}

func main() {

	go listenServerUDP()
	go writeToServerUDP()

	select {}
}

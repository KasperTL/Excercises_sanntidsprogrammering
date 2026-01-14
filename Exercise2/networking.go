package main

import (
	"fmt"
	"net"
	"runtime"
)

func writeToServerUDP() {
	raddr, err := net.ResolveUDPAddr("udp", "10.100.23.11:20017")
	//resolve a string network address into a *netUDPAddr structure
	//requireed by net.DialAddr and net.ListenUDP to establish UDP connection

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	message := []byte("Test group 17")
	conn.Write(message)

	defer conn.Close()
}

func readfromServerUDP() {
	raddr, err := net.ResolveUDPAddr("udp", ":20017")

	msg, err := net.ListenUDP("udp", raddr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	buffer := make([]byte, 1024)

	msg.Read(buffer[0:])

	fmt.Print("From server: ", string(buffer[0:]))
	defer msg.Close()
}


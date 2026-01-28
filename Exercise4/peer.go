package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os/exec"
	"time"
)

const (
	port         = 9999                   //UDP port for hearbeat comm
	timeout      = 2 * time.Second        //how long backup waits before deciding primary is dead
	tickInterval = 500 * time.Millisecond //how often primary broadcasts heartbeat
)

var counter uint64

func s_Backup() {
	// Linux (gnome-terminal)
	//err := exec.Command("gnome-terminal", "--", "go", "run", "peer.go").Run()

	//windows version
	err := exec.Command("cmd", "/C", "start", "powershell", "go", "run", "peer.go").Run()
	if err != nil {
		log.Println("Failed to spawn backup:", err)
	} else {
		fmt.Println("New backup spawned")
	}
}

func main() {

	//resolve UDP
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		log.Fatal("Failed to resolve address:", err)
	}

	// Try to listen as backup
	//if primary is alive, this fails adn exits
	//if primary is dead, this goes through and becomes new primary
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal("Failed to start as backup:", err)
	}

	log.Println("Starting as BACKUP, listening for primary...")

	//backup
	isPrimary := false
	buf := make([]byte, 8) //buffer to hold counter

	for !isPrimary {

		//if data is not received within timeout, primary is assumed dead
		conn.SetReadDeadline(time.Now().Add(timeout))
		n, _, err := conn.ReadFromUDP(buf)

		if err != nil {
			// Timeout - primary is dead!
			isPrimary = true
			log.Println("Primary timeout detected!")
		} else if n == 8 {
			// Received heartbeat with counter value
			counter = binary.BigEndian.Uint64(buf)
		}
	}

	// Close the listening socket so the next backup can use it
	conn.Close()

	//new primary?
	fmt.Println("\n" + "=========================================")
	fmt.Println("I AM NOW THE PRIMARY!")
	fmt.Println("=======================================" + "\n")
	fmt.Printf("Continuing from counter = %d\n\n", counter) //resumes from where primary left off

	// Spawn new backup BEFORE starting to broadcast
	s_Backup()
	time.Sleep(200 * time.Millisecond) // Give backup time to start listening

	// Setup UDP connection for broadcasting
	bcastConn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("Failed to create broadcast connection:", err)
	}
	defer bcastConn.Close()

	//primary
	for {
		// Print current counter value
		fmt.Printf("\tCount: %d\n", counter)

		// Broadcast current counter to backup
		//convert 64bit number into  8 bytes and store in byte array (buf)
		//cant send 64 bits in network only in bytes
		//Big endian so biggest number is at the end (ascending order) for count up
		binary.BigEndian.PutUint64(buf, counter)
		//send backup to UDP
		_, err := bcastConn.Write(buf)
		if err != nil {
			log.Println("Warning: Failed to broadcast:", err)
		}

		// Increment for next iteration
		counter++

		time.Sleep(tickInterval)
	}
}

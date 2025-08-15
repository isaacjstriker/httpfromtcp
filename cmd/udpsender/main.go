package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"log"
)

func main() {
	serverAddr := "localhost:42069"

	udpAddr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		log.Fatalf("error resolving udp address: %v", err)
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatalf("error dialing udp: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Sending to %s. Type your message and press Enter to send. Press Ctrl+C to exit.\n", serverAddr)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("error reading input: %v\n", err)
			os.Exit(1)
		}

		_, err = conn.Write([]byte(message))
		if err != nil {
			log.Fatalf("error sending message: %v", err)
			os.Exit(1)
		}

		fmt.Printf("Message sent: %s", message)
	}
}
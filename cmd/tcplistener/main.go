package main

import (
	"fmt"
	"io"
	"log"
	"strings"
	"net"
)

func getLinesChannel (conn net.Conn) <-chan string {
	lines := make(chan string)
	buffer := make([]byte, 1024)
	currentLine := ""

	go func(){
		defer conn.Close()
		defer close(lines)
		for {
			n, err := conn.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error reading connection: %v\n", err)
				return
			}

			parts := strings.Split(string(buffer[:n]), "\n")
			// For each part, print a line to the console (except the last one)
			for i := 0; i < len(parts) - 1; i++ {
					lines <- currentLine + parts[i]
					currentLine = ""
			}
			currentLine += parts[len(parts)-1]
		}
		if currentLine != "" {
			lines <- currentLine
		}
	}()
	return lines
}

func main() {
	l, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println("connection accepted")

		lines := getLinesChannel(conn)
		for line := range lines {
			fmt.Printf("%s\r\n", line)
		}
		fmt.Println("connection closed")
	}
}
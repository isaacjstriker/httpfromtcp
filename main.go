package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func getLinesChannel (f io.ReadCloser) <-chan string {
	lines := make(chan string)

	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	buffer := make([]byte, 8)
	currentLine := ""

	go func(){
		defer file.Close()
		defer close(lines)
		for {
			n, err := file.Read(buffer)
			if err == io.EOF {
				// End of file reached
				break
			}
			if err != nil {
				fmt.Printf("error reading file: %v\n", err)
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
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	lines := getLinesChannel(file)
	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}
}
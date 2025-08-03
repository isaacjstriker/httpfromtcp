package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Open messages.txt for reading
	file, err := os.Open("messages.txt")
	if err != nil {
		err = fmt.Errorf("could not open file: %w", err)
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	buffer := make([]byte, 8)

	currentLine := ""

	// Loop to read every 8 bytes
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
				fmt.Printf("read: %s\n", currentLine + parts[i])
				currentLine = ""
		}
		currentLine += parts[len(parts)-1]
	}

	if currentLine != "" {
		fmt.Printf("read: %s\n", currentLine)
	}
}
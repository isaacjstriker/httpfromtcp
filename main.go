package main

import (
	"fmt"
	"io"
	"os"
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

	// A slice holding 8 bytes of data from the file
	buffer := make([]byte, 8)

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

		// Print the 8 bytes as a string
		fmt.Printf("read: %s\n", string(buffer[:n]))
	}
}
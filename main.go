package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// reads a file and stores a error if there are any
	f, err := os.Open("message.txt")
	if err != nil {
		log.Fatal("error ", "error", err) // Find the reason why does this output needs two "error"
	}

	for {
		//byte is just an alias for uint8
		data := make([]byte, 8)
		n, err := f.Read(data)
		if err != nil {
			break
		}

		fmt.Printf("read: %s\n", string(data[:n]))
	}
}

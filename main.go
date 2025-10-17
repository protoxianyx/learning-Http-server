package main

import (
	"bytes"
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

	str := ""
	for {
		//byte is just an alias for uint8
		data := make([]byte, 8)
		n, err := f.Read(data)
		if err != nil {
			break
		}

		data = data[:n]
		if i := bytes.IndexByte(data, '\n'); i != -1 {
			str += string(data[:i])
			data = data[i+1:]

			fmt.Printf("read: %s\n", str)
			str = ""
		}

		str += string(data)
	}

	if len(str) != 0 {
		fmt.Printf("read: %s\n", str)

	}
}

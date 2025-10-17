package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func getLinesChannel(f io.ReadCloser) <-chan string { //  very weird syntax, gotta investigate
	out := make(chan string, 1)

	go func() {
		defer f.Close()
		defer close(out)

		str := "" // first stream of data that is t-c |
		for {
			//byte is just an alias for uint8
			data := make([]byte, 8)
			n, err := f.Read(data)
			if err != nil {
				break
			}

			data = data[:n] // Again stores the value of the data from the start to the 8thth position. Which is T -> c | c to the 8th position upto the moto part whrer \n i sencountered
			// fmt.Printf("data data: %s\n", data)
			if i := bytes.IndexByte(data, '\n'); i != -1 { // skips this until matches the condition aka it shoudl end in the new line. | encounetrs \n
				str += string(data[:i]) //  at this poitn the str is the whole first line, adn the I position is the end of fthe new line
				data = data[i+1:]

				out <- str
				// fmt.Printf("index: %d \n", i)
				str = ""
			}

			str += string(data) // saves upto t - c in str before the start of the for loop. t-c |
			// fmt.Printf("str data: %s\n", str)
		}

		if len(str) != 0 {
			out <- str
		}
	}()

	return out
}

func main() {

	// reads a file and stores a error if there are any
	f, err := os.Open("message.txt")
	if err != nil {
		log.Fatal("error ", "error", err) // Find the reason why does this output needs two "error"
	}

	lines := getLinesChannel(f)

	for line := range lines {
		
		fmt.Printf("second read: %s\n", line)

	}

}

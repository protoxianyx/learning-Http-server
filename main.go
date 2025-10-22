package main

import (
	"fmt"
	"log"
	"net"

	"demoproject/internal/request"
)

func main() {

	// reads a file and stores a error if there are any
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal("error ", "error", err) // Find the reason why does this output needs two "error"
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("error ", "error", err) // Find the reason why does this output needs two "error"
		}

		r, err := request.RequestFromReader(conn)
		if err != nil {
			log.Fatal("error", "error", err)
		}

		fmt.Printf("RequestLine:\n")
		fmt.Printf("- Methord: %s\n", r.RequestLine.Method)
		fmt.Printf("- Request Target: %s\n", r.RequestLine.RequestTarget)
		fmt.Printf("- Http Version: %s\n", r.RequestLine.HttpVersion)
		fmt.Printf("Http Version:\n")
		r.Headers.ForEach(func(n, v string){
			fmt.Printf("- %s: %s\n", n, v)
		})
	}

}

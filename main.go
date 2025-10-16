package main

import (
	"fmt"
	"net"
)

func main() {
	ln, _ := net.Listen("tcp", ":8080")
	defer ln.Close()

	for {
		conn, _ := ln.Accept() // accept connection
		go func(c net.Conn) {
			defer c.Close()
			fmt.Fprintf(c, "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\nHello World!")
		}(conn)
	}

}

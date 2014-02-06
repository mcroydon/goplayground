// An implementation of the QOTD protocol based on RFC 865.

package main

import (
	"log"
	"net"
)

var (
	quote = "\"In the immortal words of Jean Paul Sartre, 'Au revoir, gopher'.\"\n"
)

func qotd(conn net.Conn) {
	defer conn.Close()
	_, err := conn.Write([]byte(quote))
	if err != nil {
		log.Printf("Error during write: %v", err)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":17")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("goqotd listening for connections on port 17.")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error during accept: %v", err)
		}
		go qotd(conn)
	}
}

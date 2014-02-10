// An implementation of the QOTD protocol based on RFC 865.

package main

import (
	"log"
	"net"
)

var (
	quote = "\"In the immortal words of Jean Paul Sartre, 'Au revoir, gopher'.\"\n"
)

func handleTCP(conn net.Conn) {
	defer conn.Close()
	_, err := conn.Write([]byte(quote))
	if err != nil {
		log.Printf("Error during write: %v", err)
	}
}

func QotdTCP() {
	listener, err := net.Listen("tcp", ":17")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error during accept: %v", err)
		}
		go handleTCP(conn)
	}
}

func handleUDP(conn *net.UDPConn) {
	b := make([]byte, 1)
	_, addr, err := conn.ReadFromUDP(b)
	if err != nil {
		log.Printf("Error during accept: %v", err)
	}
	_, err = conn.WriteToUDP([]byte(quote), addr)
	if err != nil {
		log.Printf("Error during accept: %v", err)
	}
}

func QotdUDP() {
	addr, err := net.ResolveUDPAddr("udp", ":17")
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		handleUDP(listener)
	}
}

func main() {
	go QotdTCP()
	go QotdUDP()
	log.Println("goqotd listening for connections on port 17.")
	select {

	}
}

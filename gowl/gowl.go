package main

import (
	"log"
	"math/rand"
	"net"
	"time"
)

var (
	owls = []string{
		"^(OvO)^", // Sean Bleier
		`  ___
 (o,o)
<  .  >
--"-"---`, // http://www.asciiworld.com/-Owls-.html

	}
	owlCount = len(owls)
)

func check(err error, message string) {
	if err != nil {
		log.Print(message)
		panic(err)
	}
}

func owl(conn net.Conn) {
	defer conn.Close()
	var index = rand.Intn(owlCount)
	log.Printf("Serving owl %v to %v", index, conn.RemoteAddr().String())
	conn.Write([]byte(owls[index] + "\n"))
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	listener, err := net.Listen("tcp", ":3400")
	defer listener.Close()
	check(err, "Unable to listen on :3400.")
	log.Print("Serving owls on :3400.")
	for {
		conn, err := listener.Accept()
		check(err, "Unable to accept connection.")
		go owl(conn)
	}
}

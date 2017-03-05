package main

import "bytes"
import "encoding/binary"
import "flag"
import "fmt"
import "log"
import "math/rand"
import "net"
import "time"

func getAddr(ip string, port uint) string {
	return fmt.Sprintf("%v:%v", ip, port)
}

func numbercast(ip string, port uint) {
	number := rand.Uint64()
	log.Printf("now serving: %v", number)

	multicastAddr := getAddr(ip, port)

	addr, err := net.ResolveUDPAddr("udp", multicastAddr)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, number)
	if err != nil {
		log.Fatal(err)
	}
	conn.Write([]byte(buf.Bytes()))
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ip := flag.String("ip", "224.0.0.1", "IP to send UDP traffic to.")
	port := flag.Uint("port", 4815, "Port number to use.")
	seconds := flag.Uint("seconds", 1, "Number of seconds betweeen broadcast.")

	flag.Parse()

	log.Printf("Sending numbers to %v every %v seconds.", getAddr(*ip, *port), *seconds)

	ticker := time.Tick(time.Duration(*seconds) * time.Second)

	for _ = range ticker {
		numbercast(*ip, *port)
	}
}

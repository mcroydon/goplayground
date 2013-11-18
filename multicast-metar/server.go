package main

import (
	"bufio"
	"log"
	"net"
	"net/http"
	"time"
)

var (
	ipv4mcastaddr = &net.UDPAddr{
		IP:   net.ParseIP("224.0.0.251"),
		Port: 5522,
	}
	metarUrl = "http://weather.noaa.gov/pub/data/observations/metar/stations/KLWC.TXT"
)

func check(err error, message string) {
	if err != nil {
		log.Print(message)
		panic(err)
	}
}

func getMetar() string {
	resp, err := http.Get(metarUrl)
	check(err, "Error downloading METAR.")
	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)
	line := ""
	for scanner.Scan() {
		line = scanner.Text()
	}
	return line
}

func main() {
	conn, err := net.ListenMulticastUDP("udp4", nil, ipv4mcastaddr)
	defer conn.Close()
	check(err, "Error in main.")
	ticker := time.NewTicker(time.Second * 30)
	go func() {
		for t := range ticker.C {
			line := getMetar()
			log.Printf("Sending \"%s\" at %v", line, t)
			_, err = conn.WriteToUDP([]byte(line), ipv4mcastaddr)
			check(err, "Error while sending.")
		}
	}()
	select {}
}

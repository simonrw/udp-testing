package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	hostStrArg := flag.String("host", "", "host to connect to")
	flag.Parse()

	host := ":1053"
	if *hostStrArg != "" {
		host = *hostStrArg
	}

	udpServer, err := net.ListenPacket("udp", host)
	if err != nil {
		panic(err)
	}
	defer udpServer.Close()

	for {
		buf := make([]byte, 1024)
		_, addr, err := udpServer.ReadFrom(buf)
		if err != nil {
			continue
		}
		go response(udpServer, addr, buf)
	}
}

func response(udpServer net.PacketConn, addr net.Addr, buf []byte) {
	log.Println("received packet")
	time := time.Now().Format(time.ANSIC)
	responseStr := fmt.Sprintf("time received: %v your message %v", time, string(buf))
	udpServer.WriteTo([]byte(responseStr), addr)
	log.Println("response sent")
}

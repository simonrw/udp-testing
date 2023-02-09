package main

import (
	"flag"
	"net"
)

func main() {
	hostStrArg := flag.String("host", "", "host to connect to")
	flag.Parse()

	host := ":1053"
	if *hostStrArg != "" {
		host = *hostStrArg
	}

	udpServer, err := net.ResolveUDPAddr("udp", host)
	if err != nil {
		panic(err)
	}

	con, err := net.DialUDP("udp", nil, udpServer)
	if err != nil {
		panic(err)
	}
	defer con.Close()

	if _, err := con.Write([]byte("message")); err != nil {
		panic(err)
	}

	received := make([]byte, 1024)
	n, err := con.Read(received)
	if err != nil {
		panic(err)
	}
	println(string(received[:n]))
}

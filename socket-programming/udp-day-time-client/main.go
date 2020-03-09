package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	server := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp", server)
	checkError("udp address resolve failed", err)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError("udp conn failed", err)

	_, err = conn.Write([]byte("tell me the time!"))
	checkError("udp msg send failed", err)

	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError("udp read failed", err)

	fmt.Println(string(buf[0:n]))

	os.Exit(0)
}

func checkError(msg string, err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "%s: %s\n", msg, err.Error())
	os.Exit(1)
}

package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	addr := "localhost:3333"
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	checkError("udp address resolve failed", err)

	conn, err := net.ListenUDP("udp", udpAddr)
	checkError("udp listen failed", err)

	for {
		func(conn *net.UDPConn) {
			var buf [512]byte
			_, clientAddr, err := conn.ReadFromUDP(buf[0:])
			checkError("udp client read failed", err)

			now := time.Now().String()
			_, err = conn.WriteToUDP([]byte(now), clientAddr)
			checkError("udp client write failed", err)

		}(conn)
	}
}

func checkError(msg string, err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "%s: %s\n", msg, err.Error())
	os.Exit(1)
}

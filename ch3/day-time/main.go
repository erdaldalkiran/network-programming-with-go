package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":3333"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Fprintf(os.Stderr, "conn error: %s\n", err.Error())
			continue
		}

		go func(conn *net.TCPConn) {
			defer conn.Close()
			now := time.Now().String()
			fmt.Fprintf(os.Stdout, "now: %s\n", now)
			conn.Write([]byte(now))
		}(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	DIR = "DIR"
	CD  = "CD"
	PWD = "PWD"
)

func main() {
	address := "0.0.0.0:4444"
	tcpAddress, err := net.ResolveTCPAddr("tcp", address)
	checkError("resolve tcp address failed", err)

	listener, err := net.ListenTCP("tcp", tcpAddress)
	checkError("listen tcp failed", err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s", "error while accepting connection", err.Error())
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", "error while reading data from connection", err.Error())
			break
		}

		line := string(buf[0:n])
		strs := strings.SplitN(line, " ", 2)

		switch strs[0] {
		case CD:
			chDir(conn, strs[1])
		case DIR:
			listDir(conn)
		case PWD:
			pwd(conn)
		default:
			fmt.Fprintf(os.Stderr, "%s: %s", "unknown command", line)
			conn.Write([]byte("unknown command\n"))
		}
	}
}

func chDir(conn net.Conn, dir string) {
	err := os.Chdir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s", "error while changing directory", err.Error())
		conn.Write([]byte("ERROR"))
		return
	}

	conn.Write([]byte("OK"))
}

func listDir(conn net.Conn) {
	defer conn.Write([]byte("\n"))

	dir, err := os.Open(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s", "error while opening directory", err.Error())
		conn.Write([]byte("ERROR"))
		return
	}

	fileNames, err := dir.Readdirnames(-1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s", "error while reading directory contents", err.Error())
		conn.Write([]byte("ERROR"))
		return
	}

	for _, fileName := range fileNames {
		conn.Write([]byte(fileName + "\n"))
	}
}

func pwd(conn net.Conn) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s", "error while getting working directory", err.Error())
		conn.Write([]byte("ERROR"))
		return
	}

	conn.Write([]byte(dir))
}

func checkError(msg string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s", msg, err.Error())
		os.Exit(1)
	}
}

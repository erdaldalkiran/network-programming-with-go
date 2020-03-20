package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/websocket"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage", os.Args[0], "ws://host:port")
		os.Exit(1)
	}

	sAdd := os.Args[1]
	conn, err := websocket.Dial(sAdd, "", "http://origin")
	if err != nil {
		fmt.Println("unable to connect server", err.Error())
		os.Exit(1)
	}

	for {
		var msg string
		err = websocket.Message.Receive(conn, &msg)
		if err != nil {
			if err == io.EOF {
				fmt.Println("server closed the connection.")
				break
			}

			fmt.Println("receive failed", err.Error())
			os.Exit(1)
		}
		fmt.Println("received msg", msg)

		err = websocket.Message.Send(conn, msg)
		if err != nil {
			fmt.Println("send failed", err.Error())
			os.Exit(1)
		}
		fmt.Println("send msg", msg)
	}

}

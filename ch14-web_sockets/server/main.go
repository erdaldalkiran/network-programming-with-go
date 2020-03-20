package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func Echo(ws *websocket.Conn) {
	defer ws.Close()
	fmt.Println("Echoing")

	for i := 0; i < 10; i++ {
		sendMsg := fmt.Sprintf("echo%d", i)
		err := websocket.Message.Send(ws, sendMsg)
		if err != nil {
			fmt.Println("send message failed", err.Error())
			break
		}
		fmt.Println("message send:", sendMsg)

		var receiveMsg string
		err = websocket.Message.Receive(ws, &receiveMsg)
		if err != nil {
			fmt.Println("send message failed", err.Error())
			break
		}
		fmt.Println("message received:", sendMsg)
	}
}

func main() {
	http.Handle("/", websocket.Handler(Echo))
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

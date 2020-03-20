package main

import (
	"fmt"
	"net/http"
	"os"

	"ciko.com/codec/xmlcodec"
	"golang.org/x/net/websocket"
)

type Person struct {
	Name   string
	Emails []string
}

func ReceivePerson(ws *websocket.Conn) {
	var person Person
	err := xmlcodec.XMLCodec.Receive(ws, &person)
	if err != nil {
		fmt.Println("Can't receive")
	} else {

		fmt.Println("Name: " + person.Name)
		for _, e := range person.Emails {
			fmt.Println("An email: " + e)
		}
	}
}

func main() {

	http.Handle("/", websocket.Handler(ReceivePerson))
	err := http.ListenAndServe(":1234", nil)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

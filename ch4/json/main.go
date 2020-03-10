package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  Name
	Email []Email
}

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string
	Address string
}

func main() {
	var hede Person
	fmt.Println(hede)
	p := Person{
		Name: Name{Family: "Tanesi", Personal: "Ciko"},
		Email: []Email{
			{Kind: "personal", Address: "ciko@tanesi.com"},
			{"work", "tanesi@ciko.com"}}}
	writeToFile(p)
	fmt.Println("data is encoded to file")

	pr := readFromFile()
	fmt.Println(pr)
}

func writeToFile(p Person) {
	f, err := os.Create("person.json")
	checkError("file creation failed", err)
	e := json.NewEncoder(f)
	err = e.Encode(p)
	checkError("encode failed", err)
}

func readFromFile() Person {
	f, err := os.Open("person.json")
	checkError("file read failed", err)

	d := json.NewDecoder(f)
	var p Person
	err = d.Decode(&p)
	checkError("decode failed", err)
	return p
}

func checkError(msg string, err error) {
	if err != nil {

		fmt.Fprintf(os.Stderr, "%s: %s\n", msg, err.Error())
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server")
		os.Exit(1)
	}
	serverAddress := os.Args[1]

	client, err := jsonrpc.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatal("unable to dial rpc server", err.Error())
	}
	defer client.Close()

	multiply := &Args{3, 5}
	mresult := new(int)
	err = client.Call("Arith.Multiply", multiply, mresult)
	if err != nil {
		fmt.Println("multiply failed", err.Error())
	}
	fmt.Println("multiplication result", *mresult)

	division := &Args{23, 4}
	dresult := new(Quotient)
	err = client.Call("Arith.Divide", division, dresult)
	if err != nil {
		fmt.Println("division failed", err.Error())
	}
	fmt.Println("division result", *dresult)

}

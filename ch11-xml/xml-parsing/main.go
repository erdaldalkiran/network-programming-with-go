package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

func main() {
	x := `<person><name><family> Newmarch </family><personal> Jan </personal></name><email type="personal"> jan@newmarch.name </email><email type="work"> j.newmarch@boxhill.edu.au </email></person>`

	r := strings.NewReader(x)

	parser := xml.NewDecoder(r)
	depth := 0
	for {
		token, err := parser.Token()
		if err != nil {
			break
		}
		switch t := token.(type) {
		case xml.StartElement:
			name := t.Name.Local
			printElmt(name, depth)
			depth++
		case xml.EndElement:
			depth--
			name := t.Name.Local
			printElmt(name, depth)
		case xml.CharData:
			printElmt("\""+string([]byte(t))+"\"", depth)
		case xml.Comment:
			printElmt("Comment", depth)
		case xml.ProcInst:
			printElmt("ProcInst", depth)
		case xml.Directive:
			printElmt("Directive", depth)
		default:
			fmt.Println("Unknown")
		}
	}
}

func printElmt(s string, depth int) {
	for n := 0; n < depth; n++ {
		fmt.Print("  ")
	}
	fmt.Println(s)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

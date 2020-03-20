package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage", os.Args[0], "http://host:port/page")
		os.Exit(1)
	}

	url, err := url.Parse(os.Args[1])
	checkError("parsing url failed", err)

	request, err := http.NewRequest("GET", url.String(), nil)
	request.Header.Add("Accept-Charset", "UTF-8;q=1, ISO-8859-1;q=0")
	checkError("creating request failed", err)

	client := &http.Client{}

	response, err := client.Do(request)
	checkError("getting response failed", err)
	if response.StatusCode != 200 {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	var buf [512]byte
	reader := response.Body
	for {
		n, err := reader.Read(buf[0:])
		if err == io.EOF {
			break
		}
		checkError("reading response failed", err)
		fmt.Print(string(buf[:n]))
	}

}

func checkError(msg string, err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "%s: %s", msg, err.Error())
	os.Exit(1)
}

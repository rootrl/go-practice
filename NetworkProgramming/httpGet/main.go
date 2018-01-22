package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
_	"strings"
)

func main() {
	url := os.Args[1]

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	if response.Status != "200 OK" {
		fmt.Println(response.Status)

		os.Exit(0)

	}

	b, _ := httputil.DumpResponse(response, false)

	fmt.Println(string(b))

	var buf [512]byte

	reader := response.Body
	count := 0

	for {
		// 阻塞读？
		n, err := reader.Read(buf[0:])
		if err != nil {
			break
		}
		// fmt.Println("readed 512 kb \n", n, "\n")
		fmt.Println(string(buf[0:n]))
		count++	
	}

	fmt.Println(count)
	os.Exit(0)
}

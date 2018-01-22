package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	url := os.Args[1]

	response, err := http.Head(url)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.Status)

	for k, v := range response.Header {

		fmt.Println(k, ":", v)

	}

}

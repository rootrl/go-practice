package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	host := os.Args[1]

	addrs, err := net.LookupHost(host)

	if err != nil {

		panic(err)
	}

	for k, v := range addrs {
		fmt.Println(k, v)
	}

}

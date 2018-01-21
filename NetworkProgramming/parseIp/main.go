package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	// 0:0:0:0:0:0:0:1 || 127.0.0.1
	ipString := os.Args[1]

	ip := net.ParseIP(ipString)

	fmt.Println(ip)

}

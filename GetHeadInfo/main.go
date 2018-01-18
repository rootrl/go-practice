package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {

	service := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {

		panic(err)
	}

	_, err = conn.Write([]byte("HEAD / HTTP/1.1\r\n\r\n"))

	result, err := ioutil.ReadAll(conn)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))

	os.Exit(0)

}

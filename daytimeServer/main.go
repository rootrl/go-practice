package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	check(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	check(err)

	for {
		conn, err := listener.Accept()

		if err != nil {
			continue
		}

		dayTime := time.Now().String()
		conn.Write([]byte(dayTime))
		conn.Close()

	}

}

func check(err error) {

	if err != nil {

		fmt.Println("Err:", err.Error())
		os.Exit(1)

	}

}

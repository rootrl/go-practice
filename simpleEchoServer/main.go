package main

import (
	"fmt"

	"net"
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

		handleClient(conn)
		conn.Close()

	}

}

func handleClient(conn net.Conn) {

	var buf [512]byte

	for {
		n, err := conn.Read(buf[0:])

		if err != nil {
			return
		}

		fmt.Println("from client:", string(buf[0:]))

		content := "from server:" + string(buf[0:n])

		_, err2 := conn.Write([]byte(content))

		if err2 != nil {

			return
		}

	}

}

func check(err error) {

	if err != nil {
		panic(err)
	}
}

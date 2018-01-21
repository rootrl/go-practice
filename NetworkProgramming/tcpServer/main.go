package main

import (
	"net"
	"time"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1200")
	check(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	check(err)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			// log
			continue
		}

		go handlerConn(conn)

	}

}

func handlerConn(conn net.Conn) {
	time := time.Now().String()
	conn.Write([]byte(time))
	conn.Close()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

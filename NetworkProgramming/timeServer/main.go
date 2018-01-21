package main

import (
	"encoding/asn1"
	"net"
	"time"
)

func check(err error) {

	if err != nil {
		panic(err)
	}

}

func main() {

	// tcpAddr, err := net.ResolveTCPAddr("tcp", ":1201")
	// check(err)

	listener, err := net.Listen("tcp", ":1200")
	check(err)
	for {
		conn, err := listener.Accept()

		if err != nil {
			// log
			continue
		}

		time := time.Now()
		packTime, _ := asn1.Marshal(time)
		conn.Write(packTime)
		conn.Close()
	}

}

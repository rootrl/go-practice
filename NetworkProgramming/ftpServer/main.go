package main

import (
	"net"
	"os"
)

const (
	DIR = "DIR"
	CD  = "CD"
	PWD = "PWD"
)

func check(err error) {

	if err != nil {
		panic(err)
	}

}

func main() {

	listener, err := net.Listen("tcp", ":1200")

	check(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {

	defer conn.Close()

	var buf [512]byte

	for {

		n, err := conn.Read(buf[0:])
		if err != nil {
			conn.Close()
			return
		}

		s := string(buf[0:n])

		if s[0:2] == CD {
			chdir(conn, s[3:])
		} else if s[0:3] == DIR {

			dirList(conn)
		} else if s[0:3] == PWD {
			pwd(conn)
		} else {
			conn.Write([]byte("surpport: cmd\\pwd\\dir \r\n"))
		}
	}

}

func chdir(conn net.Conn, s string) {

	if os.Chdir(s) == nil {
		conn.Write([]byte("OK"))
	} else {
		conn.Write([]byte("error"))
	}

}

func pwd(conn net.Conn) {
	s, err := os.Getwd()
	if err != nil {
		conn.Write([]byte(""))
	} else {
		conn.Write([]byte(s))
	}

}

func dirList(conn net.Conn) {

	defer conn.Write([]byte("\r\n"))

	dir, err := os.Open(".")
	if err != nil {
		return
	}

	names, err := dir.Readdirnames(-1)

	if err != nil {
		return
	}

	for _, nm := range names {
		conn.Write([]byte(nm + "\r\n"))
	}

}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	uiDir  = "dir"
	uiCd   = "cd"
	uiPwd  = "pwd"
	uiQuit = "quit"
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

	conn, err := net.Dial("tcp", ":1200")
	
	check(err)	

	reader := bufio.NewReader(os.Stdin)

	for {

		line, err := reader.ReadString('\n')

		line = strings.TrimRight(line, "\t\r\n")

		if err != nil {
			break
		}

		strs := strings.SplitN(line, " ", 2)

		switch strs[0] {
		case uiDir:
			dirRequest(conn)
		case uiCd:
			if len(strs) != 2 {
				fmt.Println("cd <dir>")
				continue
			}

			fmt.Println("cd \"", strs[1], "\"")
			cdRequest(conn, strs[1])

		case uiPwd:
			pwdRequest(conn)

		case uiQuit:
			conn.Close()
			os.Exit(0)
		default:
			fmt.Println("unknow command")
		}

	}

}

func dirRequest(conn net.Conn) {
	conn.Write([]byte(DIR + " "))

	var buf [512]byte

	result := bytes.NewBuffer(nil)

	for {
		n, _ := conn.Read(buf[0:])
		result.Write(buf[0:n])
		length := result.Len()

		contents := result.Bytes()
		if string(contents[length-4:]) == "\r\n\r\n" {
			fmt.Println(string(contents[0 : length-4]))
			return
		}
	}

}

func cdRequest(conn net.Conn, dir string) {
	conn.Write([]byte(CD + " " + dir))
	var response [512]byte

	n, _ := conn.Read(response[0:])

	s := string(response[0:n])

	if s != "OK" {
		fmt.Println("Faild to change dir")
	}

}

func pwdRequest(conn net.Conn) {
	conn.Write([]byte(PWD))
	var response [512]byte
	n, _ := conn.Read(response[0:])
	s := string(response[0:n])
	fmt.Println("current dir \"" + s + "\"")

}

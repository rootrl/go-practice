package main

import (
	"encoding/asn1"
	"fmt"
	"io"
	"net"
	"time"
	"bytes"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:1200")
	check(err)

	result, err := readFully(conn)
	var newTime time.Time
	_, err1 := asn1.Unmarshal(result, &newTime)
	check(err1)

	fmt.Println(newTime.String())

}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte

	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}

	}

	return result.Bytes(), nil

}

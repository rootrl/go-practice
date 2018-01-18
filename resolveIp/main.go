package main


import (
"fmt"
"os"
"net"

)


func main() {

if len(os.Args) < 2 {
fmt.Println("usage: ./resolveIp www.baidu.com")
os.Exit(1)
}

url := os.Args[1]


addr, err := net.ResolveIPAddr("ip", url)

if err != nil {

panic(err)

}

fmt.Println(addr.String())

}

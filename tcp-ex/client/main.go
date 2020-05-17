package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 2020,
	})
	defer conn.Close()
	if err != nil {
		return
	}
	n, err := conn.Write([]byte("hello"))
	if err != nil {
		return
	}
	fmt.Println("send ok", n)
}

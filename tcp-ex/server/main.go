package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 2020,
	})
	defer listener.Close()

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			return
		}
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		if err != nil {
			return
		}
		fmt.Println(string(data[:n]))
	}
}

package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9090,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		data := make([]byte, 1024)
		// 从客户端读取消息
		n, addr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		go func() {
			fmt.Println(addr, string(data[:n]), n)
			// 向客户端发送消息
			_, err = conn.WriteToUDP(data[:n], addr)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}
}

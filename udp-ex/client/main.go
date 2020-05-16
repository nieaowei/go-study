package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 9090,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 100; i++ {
		// 向服务器发送消息
		_, err = conn.Write([]byte("hello"))
		if err != nil {
			fmt.Println(err)
			return
		}
		// 从服务器读取消息
		res := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(res)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(addr, n, string(res[:n]))
	}
}

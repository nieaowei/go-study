/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/27 3:11 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("start client...")
	listen, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println("start listen faild...")
		return
	}
	for {
		coon, err := listen.Accept()
		if err != nil {
			fmt.Println("accept faild...")
			continue
		}
		go func() {
			defer coon.Close()
			for {
				buf := make([]byte, 512)
				_, err := coon.Read(buf)
				if err != nil {
					fmt.Println("read err :", err)
					return
				}
				fmt.Println("read:", string(buf))
			}
		}()
	}
}

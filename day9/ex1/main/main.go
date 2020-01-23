/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/27 3:04 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("start server...")
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("dial faild...", err)
		return
	}
	defer conn.Close()
	input := bufio.NewReader(os.Stdin)
	for {
		line, err := input.ReadString('\n')
		if err != nil {
			fmt.Println("read filad")
		}
		trimeInput := strings.Trim(line, "\r\n")
		_, err = conn.Write([]byte(trimeInput))
		if err != nil {
			return
		}
	}
}

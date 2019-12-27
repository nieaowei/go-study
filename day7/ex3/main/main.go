/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/25 4:22 下午
 *  Notes       :	终端读写
 *******************************************************/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hellworld")
	str := "nieaowei"
	fmt.Fprintf(os.Stdout, str)
	fmt.Fscanln(os.Stdin, str)
	fmt.Println(str)

	file, err := os.OpenFile("1.txt", os.O_RDWR, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Println("open erro", err)
		return
	}
	len1, err := fmt.Fprintf(file, "errors1")
	if err != nil {

	}
	if len1 == len("errors1") {

	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input:")
	content, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Println("content:", content)

	}
}

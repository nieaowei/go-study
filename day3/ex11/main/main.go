/*******************************************************
 *  File :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/13 8:09 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(100)

	fmt.Println("生成的数：", n)

	nX := -1
	flag := false
	for {

		fmt.Scanf("%d", &nX)

		switch true {
		case nX > n:
			fmt.Println(">")
		case nX < n:
			fmt.Println("<")
		case nX == n:
			fmt.Println("it is right.")
			flag = true

		}
		if flag {
			break
		}
	}
}

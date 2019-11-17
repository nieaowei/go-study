/*******************************************************
 *  File :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/16 7:46 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
)

func calc(n int) int {
	if n == 1 {
		return 1
	}
	return n * calc(n-1)
}

func recusive(n int) {
	fmt.Println("hello", n)
	//time.Sleep(time.Second*1)
	//n++
	//if n > 3 {
	//	return
	//}
	recusive(n + 1)
}

func fab(n int) int {
	if n <= 1 {
		return 1
	}
	return fab(n-1) + fab(n-2)
}

func main() {
	recusive(0)
	//fmt.Println(calc(3))
	//	fmt.Println(fab(3))
}

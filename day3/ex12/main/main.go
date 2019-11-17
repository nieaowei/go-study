/*******************************************************
 *  File :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/13 8:15 下午
 *  Notes       :
 *******************************************************/
package main

import "fmt"

func printA(n int) (res bool) {
	if n < 0 {
		return
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("A")
		}
		fmt.Println()
	}
	return true
}

func main() {
	printA(6)

	str := "hellowrod中国"

	for key, value := range str {
		fmt.Println(key, value, len(string(value)))
	}
	fmt.Println(len(str))
}

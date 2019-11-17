/*******************************************************
 *  File :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/13 7:54 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "123"
	//fmt.Scanf("%s",&str)

	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)

	a := "12"
	//The "switch" doesn't need "break" in Golang.
	switch a {
	case "10":
		fmt.Println(10)
	case "12":
		fmt.Println(12)
		fmt.Println("././")
		fallthrough
	case "13":
		fmt.Println(13)
	default:
		fmt.Println(0)
	}

	b := 12
	switch b {
	case 0, 1, 2, 12:
		fmt.Println("case one")
	case 13:
		fmt.Println("case two")
	}

	switch true {
	case b > 10:
		fmt.Println(">")
	case b < 10:
		fmt.Println("<")
	}

}

/*******************************************************
 *  File :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/16 8:27 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"strings"
)

func Adder() func(int) int {
	x := 0

	return func(i int) int {
		x += i
		//fmt.Println(i,x)
		return x
	}
}

func makeSuffix(suffix string) func(name string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {

	f := Adder()

	fmt.Println(Adder()(2))
	fmt.Println(Adder()(3))

	fmt.Println(f(1))
	fmt.Println(f(2))

	suf := makeSuffix(".jpg")

	fmt.Println(suf("name"))
	fmt.Println(suf("name1.jpg"))

}

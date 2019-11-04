/*******************************************************
 *  File :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/4 5:10 下午
 *  Notes       :
 *******************************************************/
package main

import "fmt"

func main() {
	var a int
	var b bool
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%b\n", 10)
	fmt.Printf("%e\n", 10000000000000)
	fmt.Printf("%q\n", "dsada")
	fmt.Printf("%p\n", &a)
	fmt.Printf("%x", &a)

	str := fmt.Sprintf("%d", 123)
	str = fmt.Sprintln(123)
	fmt.Printf("%q", str)
}

/*******************************************************
 *  File :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/8 5:34 下午
 *  Notes       :
 *******************************************************/
package main

import "fmt"

func setN(p *int, n int) {
	*p = n
}

func main() {

	a := 1
	var p *int = &a
	fmt.Println("变量a的地址:", &a)
	fmt.Println("变量p的值:", p)
	fmt.Println("变量p指向内存的值:", *p)

	setN(&a, 10)
	fmt.Println("变量p指向内存的值:", *p)

}

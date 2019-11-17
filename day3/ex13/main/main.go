/*******************************************************
 *  File :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/13 8:29 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
)

/**
	函数
1.不支持重载
2.可以赋值给变量
3.匿名函数
4.多返回值

*/

func add(a, b int) (res int) {
	return a + b
}

func add3(p func(int, int) int, c int) (res int) {
	return p(c, 1)
}

func main() {
	//2.匿名函数
	c := add

	fmt.Println(add)
	fmt.Println(c)
	//a , b:= 10 ,12

	fmt.Println(c(1, 3))
	//3.函数参数
	fmt.Println(add3(add, 9))
	//4.匿名函数

}

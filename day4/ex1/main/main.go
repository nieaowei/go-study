/*******************************************************
 *  File :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/16 5:20 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
)

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	b := 0
	a := 100 / b
	fmt.Println(a)
	return
}

func test1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("diy error")

}

func main() {
	//test()
	test1()

	i := 10
	fmt.Println(&i)
	fmt.Println(i)
	j := new(int64)
	fmt.Println(j)
	*j = 10
	fmt.Println(*j)
	var a [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println(a)

	var b []int
	b = append(b, 10, 21, 21)
	b = append(b, b...)
	fmt.Println(b)

}

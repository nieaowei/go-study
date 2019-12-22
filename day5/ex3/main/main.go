/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/22 4:33 下午
 *  Notes       :
 *******************************************************/
package main

import "fmt"

type Test struct {
	num int
}

func Alter(test Test, num int) {
	test.num = num
}

func (p Test) ALter(num int) {
	p.num = num
}

func (p *Test) ALter1(num int) {
	p.num = num
}

func main() {
	t := Test{num: 1}
	fmt.Println(t)
	t.num = 2
	fmt.Println(t)
	Alter(t, 3)
	fmt.Println(t)
	(&t).ALter(10)
	fmt.Println(t)
	t.ALter1(12)
	fmt.Println(t)

}

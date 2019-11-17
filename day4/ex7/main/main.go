/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/17 2:24 下午
 *  Notes       :
 *******************************************************/
package main

import "fmt"

func main() {
	var a []int //=[]int{1,3,4}
	var b [2]int
	a = make([]int, 10)
	fmt.Printf("%T", a)
	a = append(a, 5)
	fmt.Println(a)
	fmt.Printf("%T", b)
	fmt.Println(b)
	testslice()
}

type slice struct {
	p   *[100]int
	len int
	cap int
}

func modify(a []int) {
	a[1] = 10
}

func makeSlice(s slice, cap int) slice {
	s.cap = cap
	s.p = new([100]int)
	s.len = 0
	return s
}

func testslice() {
	var s slice
	s = makeSlice(s, 10)
	s.p[0] = 1
	fmt.Println(s.p)

	var a []int = []int{1, 2}
	modify(a)
	fmt.Println(a)
}

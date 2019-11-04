/*******************************************************
 *  ProjectName :go-stduy
 *  Author      :nieaowei
 *  Date        :2019/11/3
 *  Notes       :The life cycle of the variable.
 *******************************************************/
package main

var a string

func main() {
	a = "g"
	println(a)
	f1()
}

func f1() {
	a := "o"
	println(a)
	f2()
}

func f2() {
	println(a)
}

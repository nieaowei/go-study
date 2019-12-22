/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/22 4:23 下午
 *  Notes       :
 *******************************************************/
package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float32
	Next  *Student
}

func main() {
	head := &Student{
		Name:  "nieaowei",
		Age:   0,
		Score: 0,
		Next:  nil,
	}
	stu1 := &Student{
		Name:  "nieaowei1",
		Age:   12,
		Score: 12,
		Next:  nil,
	}
	head.Next = stu1
	p := head
	for p != nil {
		fmt.Println(p)
		p = p.Next
	}
}

/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/23 2:33 下午
 *  Notes       :
 *******************************************************/
package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	//i:=0
	//s:="dsad"

	stu := Student{
		Name: "nieaowei",
		Age:  19,
	}
	var t interface{} = stu
	switch t.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Student:
		fmt.Println("Student")
	}
}

/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/22 4:04 下午
 *  Notes       :
 *******************************************************/
package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float32
}

func main() {
	var stu Student
	var stu1 Student = Student{
		Name:  "",
		Age:   0,
		Score: 0,
	}
	fmt.Println(stu1)
	stu.Name = "nieaowei"
	stu.Age = 20
	stu.Score = 99.9
	fmt.Println(stu)
	fmt.Printf("%p \n%p \n%p \n%p\n", &stu, &stu.Name, &stu.Age, &stu.Score)
}

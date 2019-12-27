/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/27 12:50 下午
 *  Notes       :
 *******************************************************/
package main

import "fmt"

type student struct {
	name string
}

func main() {
	stuChan := make(chan interface{}, 10)
	stu := student{name: "stud01"}
	stuChan <- &stu
	var stu01 interface{}
	stu01 = <-stuChan
	fmt.Println(stu01.(*student))

}

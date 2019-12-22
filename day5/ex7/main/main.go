/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/22 5:16 下午
 *  Notes       :
 *******************************************************/
package main

import "fmt"

type Test struct {
	name string
	age  int
}

type get interface {
	getName() string
	//getAge()
}

func (p *Test) getName() string {
	return p.name
}

func main() {
	t := &Test{
		name: "nieaowei",
		age:  123,
	}

	fmt.Println(t.getName())

	var get1 get = t
	fmt.Println(get1.getName())
}

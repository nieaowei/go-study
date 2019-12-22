/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/22 5:10 下午
 *  Notes       :
 *******************************************************/
package main

import "fmt"

type Test struct {
	Name string
	Age  int
}

func (p Test) String() string {
	return fmt.Sprintf("name:%s age:%d", p.Name, p.Age)
}

func main() {
	t := Test{
		Name: "sad",
		Age:  123,
	}
	fmt.Println(t)
}

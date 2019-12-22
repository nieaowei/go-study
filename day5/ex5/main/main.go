/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/22 4:45 下午
 *  Notes       :
 *******************************************************/
package main

import "fmt"

type Car struct {
	name string
	age  string
}

type Train struct {
	Car
	name string
}

func main() {
	t := &Train{
		Car: Car{
			"010",
			"1099",
		},
		name: "001",
	}
	fmt.Println(t.Car.name)
}

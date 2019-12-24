/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/23 2:47 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"go_study/day7/ex2/balance"
)

func main() {
	insts := make([]*balance.Instance, 10)
	for i := 0; i < 16; i++ {
		host := fmt.Sprintf("192.168.1.%d", i)
		port := i
		one := balance.NewInstance(host, port)
		insts = append(insts, one)
	}
}

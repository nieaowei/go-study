/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/27 2:48 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(time.Second)
	for v := range t.C {
		fmt.Println("hello", v)
	}
}

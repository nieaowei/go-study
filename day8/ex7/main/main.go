/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/27 2:42 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	ch1 := make(chan int, 10)

	for i := 0; i < 10; i++ {
		ch <- i
		ch1 <- i * 2
	}
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case v := <-ch1:
			fmt.Println(v)
		default:
			fmt.Println("finish")
			time.Sleep(time.Second)
		}
	}
}

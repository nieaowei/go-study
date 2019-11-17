/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/17 3:55 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"time"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	go func() {
		for i := len(a) - 1; i > -1; i-- {
			a[i] = i
			fmt.Printf("g1 %d %d \n", i, a[i])
		}
	}()
	go func() {
		i := 10
		for key, _ := range a {
			a[key] = i
			fmt.Printf("g2 %d %d \n", key, i)
			i++
		}
	}()
	time.Sleep(3 * time.Second)
	fmt.Println(a)
}

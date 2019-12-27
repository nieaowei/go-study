/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/27 2:55 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("main finish")
	var ap map[string]string
	go func() {
		defer func() {
			recover()
		}()
		ap["hellp"] = "asd"
	}()
	//recover()
	time.Sleep(time.Second * 3)
	fmt.Println("recover")
}

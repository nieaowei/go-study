/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/27 1:09 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"time"
)

func sendDataStr(ch chan interface{}, str string) {
	ch <- str
}

func readDataStr(ch chan interface{}) (res string) {
	fmt.Println((<-ch).(string))
	return "1"
}

func main() {
	ch := make(chan interface{})
	go sendDataStr(ch, "hello")
	go readDataStr(ch)
	time.Sleep(time.Second)
}

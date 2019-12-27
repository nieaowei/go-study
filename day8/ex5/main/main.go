/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/27 1:19 下午
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

func calc(data chan interface{}, rec chan interface{}) {
	for v := range data {
		flag := true
		for i := 2; i < v.(int); i++ {
			if v.(int)%i == 0 {
				flag = false
				break
			}
		}
		rec <- flag
	}
}

func main() {
	ch := make(chan interface{}, 100)
	go func() {
		for i := 0; i < 1000; i++ {
			ch <- i
		}
		close(ch)
	}()

	chRes := make(chan interface{}, 100)
	for i := 0; i < 8; i++ {
		go calc(ch, chRes)
	}
	go func() {
		for v := range chRes {
			fmt.Println(v.(bool))
		}
		close(chRes)
	}()

	time.Sleep(time.Second)
}

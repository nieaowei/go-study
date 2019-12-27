/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/27 2:29 下午
 *  Notes       :
 *******************************************************/
package main

import "fmt"

func send(ch chan int, exitCh chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	exitCh <- 1
}

func recv(ch chan int, exitCh chan int) {
	for v := range ch {
		fmt.Println(v)
	}
	exitCh <- 1
}

func main() {
	data := make(chan int)
	exitCh := make(chan int)
	go send(data, exitCh)
	go recv(data, exitCh)
	toall := 0
	for _ = range exitCh {
		toall++
		if toall == 2 {
			close(exitCh)
		}
	}
}

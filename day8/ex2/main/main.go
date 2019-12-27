/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/25 9:45 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"sync"
	"time"
)

var m = make(map[int]int)

var lock sync.Mutex

type task struct {
	n int
}

func calc(t *task) {
	sum := 1
	for i := 1; i <= t.n; i++ {
		sum *= i
	}
	fmt.Println(t.n, sum)
	//lock.Lock()
	//m[t.n]=sum
	//lock.Unlock()
}

func main() {
	for i := 0; i < 10; i++ {
		t := &task{i}
		go calc(t)
	}
	time.Sleep(10 * time.Second)
	//lock.Lock()
	//for k, v := range m {
	//	fmt.Println(k,":",v)
	//}
	//lock.Unlock()
}

/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/25 7:38 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"runtime"
)

func main() {
	//go func() {
	//	for i := 0; ; i ++ {
	//		fmt.Println("go:",i)
	//		time.Sleep(time.Second)
	//	}
	//}()
	//for i := 0; ; i++ {
	//	fmt.Println("main:",i)
	//	time.Sleep(time.Second)
	//}
	fmt.Println("cpuNum:", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
}

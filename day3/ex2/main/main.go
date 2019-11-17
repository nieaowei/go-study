/*******************************************************
 *  File :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/8 2:58 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"time"
)

func main() {

	nowTime := time.Now()

	fmt.Println(nowTime.Format("2006-1-02 15:04:05"))

	startTime := time.Now().UnixNano()

	fmt.Println(startTime)

	for i := 0; i < 100; i++ {

	}

	endTime := time.Now().UnixNano()

	fmt.Println(endTime - startTime)

}

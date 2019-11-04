/*******************************************************
 *  ProjectName :go-stduy
 *  Author      :nieaowei
 *  Date        :2019/11/3
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"time"
)

const (
	Man    = 1
	Famale = 2
)

func main() {
	for {
		second := time.Now().Unix()
		if second%Famale == 0 {
			fmt.Println("female")
		} else {
			fmt.Println("man")
		}
		time.Sleep(time.Second)
	}
}

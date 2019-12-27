/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/25 5:03 下午
 *  Notes       :	自定义错误
 *******************************************************/
package main

import (
	"fmt"
	"strconv"
	"time"
)

type TestError struct {
	msg  string
	id   int
	time string
}

func (p *TestError) Error() string {
	return "[" + strconv.Itoa(p.id) + "]" + p.time + p.msg
}

var TestError1 TestError = TestError{
	msg:  "error",
	id:   1,
	time: time.Now().String(),
}

func test() error {
	return &TestError1
}

func main() {
	fmt.Println(time.Now().Format("2006"))
	err := test()
	if err != nil {
		fmt.Println()
	}
	fmt.Println(err)
}

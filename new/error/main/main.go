/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/21 5:43 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"errors"
	"fmt"
)

type DIYError struct {
	msg string
	err error
}

func (p *DIYError) Error() string {
	return p.msg
}

func (p *DIYError) UnWrap() error {
	return p.err
}

func main() {
	err1 := errors.New("new error")
	err2 := fmt.Errorf("err2:[%W]", err1)
	fmt.Println(err1)
	fmt.Println(err2)
	fmt.Println(errors.Is(err2, err1))
}

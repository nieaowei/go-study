/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2020/1/22 1:16 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"net/http"
)

var url = []string{
	"http://www.baidu.com",
	"http://www.qq.com",
	"http://google.com",
}

func main() {
	c := http.Client{
		Transport:     http.DefaultTransport,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}
	for _, v := range url {
		resp, err := c.Head(v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("succ", resp.Status, resp.StatusCode)
	}
}

/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/27 4:30 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("connect fail")
	}
	defer conn.Close()
	_, err = conn.Do("mset", "abc", 123, "ac", 456)
	if err != nil {
		fmt.Println("set faild..", err)
	}
	val, err := conn.Do("get", "abc")
	if err != nil {
		fmt.Println("get faild")
	}
	fmt.Println(string(val.([]uint8)))
	err = conn.Send("get", "ac")
	if err != nil {
		fmt.Println("get faild")
	}
	conn.Flush()
	val, err = conn.Receive()
	if err != nil {
		fmt.Println("rec faild")
	}
	fmt.Println(string(val.([]uint8)))
	_, err = conn.Do("HSet", "abcd", "b", 100)
	if err != nil {
		fmt.Println("read faild", err)
	}
	val, err = redis.String(conn.Do("HGet", "abcd", "b"))
	if err != nil {
		fmt.Println("get faild")
	}
	fmt.Println(val)
}

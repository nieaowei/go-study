/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/22 4:37 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string  `json:"name"`
	age   int     `json:"age"`
	score float32 `json:"score"`
}

func main() {
	var stu Student = Student{
		Name:  "nieaowei",
		age:   18,
		score: 19,
	}
	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
}

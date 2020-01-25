/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2020/1/24 2:43 上午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"time"
)

func main() {
	g.Cfg().SetFileName("./config.ini")
	for {
		fmt.Println(g.Cfg().GetInt("database.my"))
		time.Sleep(time.Second)
	}

}

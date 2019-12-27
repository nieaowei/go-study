/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/23 2:47 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"go_study/day7/ex2/balance"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	var insts []*balance.Instance
	for i := 0; i < 10; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		port, _ := strconv.Atoi(fmt.Sprintf("880%d", i))
		one := balance.NewInstance(host, port)
		insts = append(insts, one)
	}

	var name = "round"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	for {
		inst, err := balance.DoBalance(name, insts)
		if err != nil {
			fmt.Println("do balance err")
			time.Sleep(time.Second)
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}

}

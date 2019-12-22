/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/22 6:58 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"reflect"
)

func main() {

	a := 10
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))
}
